package account

import (
	"context"
	"encoding/json"
	"errors"
	"html"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"glintfed/ent"
	"glintfed/internal/data"
	"glintfed/internal/lib/libstr"
	"glintfed/pkg/cache"
)

const defaultAvatarPath = "/storage/avatars/default.jpg"

const (
	accountCacheKeyPrefix = "pf:services:account:"
	accountCacheTTL       = 12 * time.Hour
)

var (
	ErrProfileNotFound = errors.New("account profile not found")
	htmlTagPattern     = regexp.MustCompile(`<[^>]*>`)
)

type Profile struct {
	cfg *data.Config

	profileModel ProfileModel
}

//go:generate go tool moq -rm -out mock_profile_model.go . ProfileModel
type ProfileModel interface {
	GetAccountProfile(ctx context.Context, profileID uint64) (*ent.Profile, error)
	AccountHiddenCounts(ctx context.Context, userID uint64) (hideFollowing bool, hideFollowers bool, err error)
	IsAdminAccount(ctx context.Context, profileID uint64) (bool, error)
	AccountPronouns(ctx context.Context, profileID uint64) (*ent.UserPronoun, error)
}

func NewProfile(cfg *data.Config, profileModel ProfileModel) *Profile {
	return &Profile{
		cfg: cfg,

		profileModel: profileModel,
	}
}

type GetProfileParams struct {
	ProfileID uint64
}

type ProfileResult struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	Acct           string   `json:"acct"`
	DisplayName    *string  `json:"display_name"`
	Discoverable   bool     `json:"discoverable"`
	Locked         bool     `json:"locked"`
	FollowersCount int      `json:"followers_count"`
	FollowingCount int      `json:"following_count"`
	StatusesCount  int      `json:"statuses_count"`
	Note           string   `json:"note"`
	NoteText       *string  `json:"note_text"`
	URL            string   `json:"url"`
	Avatar         string   `json:"avatar"`
	Website        *string  `json:"website"`
	Local          bool     `json:"local"`
	IsAdmin        bool     `json:"is_admin"`
	CreatedAt      *string  `json:"created_at"`
	HeaderBG       *string  `json:"header_bg"`
	LastFetchedAt  *string  `json:"last_fetched_at"`
	Pronouns       []string `json:"pronouns"`
	Location       *string  `json:"location"`
}

func (svc *Profile) GetProfile(ctx context.Context, params GetProfileParams) (*ProfileResult, error) {
	key := accountCacheKey(params.ProfileID)
	if res, ok := cachedProfile(ctx, key); ok {
		return res, nil
	}

	var res *ProfileResult
	var loadErr error
	err := cache.SetFunc(ctx, key, func() any {
		res, loadErr = svc.loadProfile(ctx, params.ProfileID)
		if loadErr != nil {
			return loadErr
		}
		return res
	}, accountCacheTTL)
	if loadErr != nil {
		return nil, loadErr
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (svc *Profile) loadProfile(ctx context.Context, profileID uint64) (*ProfileResult, error) {
	p, err := svc.profileModel.GetAccountProfile(ctx, profileID)
	if ent.IsNotFound(err) {
		return nil, ErrProfileNotFound
	}
	if err != nil {
		return nil, err
	}
	if libstr.FromPtr(p.Status) == "delete" {
		return nil, ErrProfileNotFound
	}

	local := p.UserID != nil && p.PrivateKey != nil
	hideFollowing, hideFollowers, err := svc.hiddenCounts(ctx, p, local)
	if err != nil {
		return nil, err
	}

	pronouns, err := svc.pronouns(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	isAdmin, err := svc.isAdmin(ctx, p, local)
	if err != nil {
		return nil, err
	}

	username, acct := accountNames(libstr.FromPtr(p.Username), local)
	return &ProfileResult{
		ID:             strconv.FormatUint(p.ID, 10),
		Username:       username,
		Acct:           acct,
		DisplayName:    p.Name,
		Discoverable:   true,
		Locked:         p.IsPrivate,
		FollowersCount: visibleCount(p.FollowersCount, hideFollowers),
		FollowingCount: visibleCount(p.FollowingCount, hideFollowing),
		StatusesCount:  visibleCount(p.StatusCount, false),
		Note:           libstr.FromPtr(p.Bio),
		NoteText:       noteText(p.Bio),
		URL:            p.Url(svc.baseURL()),
		Avatar:         svc.avatarURL(p),
		Website:        p.Website,
		Local:          local,
		IsAdmin:        isAdmin,
		CreatedAt:      timeJSON(p.CreatedAt),
		HeaderBG:       p.HeaderBg,
		LastFetchedAt:  timeJSON(p.LastFetchedAt),
		Pronouns:       pronouns,
		Location:       p.Location,
	}, nil
}

func (svc *Profile) hiddenCounts(ctx context.Context, p *ent.Profile, local bool) (hideFollowing bool, hideFollowers bool, err error) {
	if !local || p.UserID == nil {
		return false, false, nil
	}
	return svc.profileModel.AccountHiddenCounts(ctx, *p.UserID)
}

func (svc *Profile) isAdmin(ctx context.Context, p *ent.Profile, local bool) (bool, error) {
	if !local {
		return false, nil
	}
	return svc.profileModel.IsAdminAccount(ctx, p.ID)
}

func (svc *Profile) pronouns(ctx context.Context, profileID uint64) ([]string, error) {
	res, err := svc.profileModel.AccountPronouns(ctx, profileID)
	if ent.IsNotFound(err) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	var pronouns []string
	if err := json.Unmarshal([]byte(libstr.FromPtr(res.Pronouns)), &pronouns); err != nil {
		return []string{}, nil
	}
	if pronouns == nil {
		return []string{}, nil
	}
	return pronouns, nil
}

func accountCacheKey(profileID uint64) string {
	return accountCacheKeyPrefix + strconv.FormatUint(profileID, 10)
}

func cachedProfile(ctx context.Context, key string) (*ProfileResult, bool) {
	val := cache.Get(ctx, key)
	switch v := val.(type) {
	case *ProfileResult:
		return v, v != nil
	case ProfileResult:
		return &v, true
	default:
		return nil, false
	}
}

func (svc *Profile) avatarURL(p *ent.Profile) string {
	if p.Edges.Avatar == nil {
		return svc.storageURL(defaultAvatarPath)
	}

	avatar := p.Edges.Avatar
	cdnURL := libstr.FromPtr(avatar.CdnURL)
	if strings.HasPrefix(cdnURL, "https://") {
		return cdnURL
	}

	path := libstr.FromPtr(avatar.MediaPath)
	if path == "" || path == "public/avatars/default.jpg" || !strings.HasPrefix(path, "public") {
		return svc.storageURL(defaultAvatarPath)
	}

	path = strings.TrimPrefix(path, "public")
	if avatar.ChangeCount > 0 {
		path += "?v=" + strconv.FormatUint(uint64(avatar.ChangeCount), 10)
	}
	return svc.storageURL("/storage" + path)
}

func (svc *Profile) baseURL() string {
	if svc.cfg != nil && svc.cfg.App.URL != nil {
		return svc.cfg.App.URL.String()
	}
	if svc.cfg != nil && svc.cfg.App.URLValue != "" {
		return svc.cfg.App.URLValue
	}
	return ""
}

func (svc *Profile) storageURL(path string) string {
	path, query, _ := strings.Cut(path, "?")
	if svc.cfg != nil && svc.cfg.App.URL != nil {
		u := svc.cfg.App.URL.JoinPath(path)
		u.RawQuery = query
		return u.String()
	}
	u, err := url.Parse(strings.TrimRight(svc.baseURL(), "/") + path)
	if err != nil {
		return strings.TrimRight(svc.baseURL(), "/") + path
	}
	u.RawQuery = query
	return u.String()
}

func accountNames(raw string, local bool) (username string, acct string) {
	if local {
		return raw, raw
	}
	acct = strings.TrimPrefix(raw, "@")
	username, _, _ = strings.Cut(acct, "@")
	return username, acct
}

func visibleCount(count *uint, hidden bool) int {
	if hidden || count == nil {
		return 0
	}
	return int(*count)
}

func noteText(bio *string) *string {
	if bio == nil {
		return nil
	}
	text := htmlTagPattern.ReplaceAllString(*bio, "")
	text = html.UnescapeString(text)
	return &text
}

func timeJSON(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.UTC().Format(time.RFC3339Nano)
	return &s
}
