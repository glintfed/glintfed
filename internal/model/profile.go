package model

import (
	"context"
	"glintfed/ent"
	"glintfed/ent/profile"
	"glintfed/ent/story"
	"glintfed/ent/user"
	"glintfed/ent/userpronoun"
	"glintfed/ent/usersetting"
	"glintfed/internal/data/client"
)

type Profile struct {
	*ent.ProfileClient
	db *ent.Client
}

func NewProfile(client *client.Database) *Profile {
	return &Profile{
		ProfileClient: client.Ent.Profile,
		db:            client.Ent,
	}
}

// GetByUsernameAndID
//
//	SELECT *
//	FROM stories
//	INNER JOIN profiles ON stories.profile_id = profiles.id
//	WHERE profiles.username = ?
//		AND profiles.domain IS NULL
//		AND stories.id = ?
//		AND stories.active = true
//	LIMIT 1
func (m *Profile) GetStory(ctx context.Context, username string, storyID uint64) (*ent.Story, error) {
	return m.Query().
		Where(
			profile.Username(username),
			profile.DomainIsNil(),
		).
		QueryStories().
		Where(
			story.ID(storyID),
			story.Active(true),
		).
		WithProfile().
		First(ctx)
}

// GetLocalByUsername
//
//	SELECT *
//	FROM profiles
//	WHERE username = ?
//		AND domain IS NULL
//		AND status IS NULL
//	LIMIT 1
func (m *Profile) GetLocalByUsername(ctx context.Context, username string) (*ent.Profile, error) {
	return m.Query().
		Where(
			profile.Username(username),
			profile.DomainIsNil(),
			profile.StatusIsNil(),
		).
		First(ctx)
}

// RemoteURLExists
//
//	SELECT EXISTS(
//	  SELECT 1 FROM profiles
//	  WHERE remote_url = ?
//	)
func (m *Profile) RemoteURLExists(ctx context.Context, remoteURL string) (bool, error) {
	return m.Query().
		Where(profile.RemoteURL(remoteURL)).
		Exist(ctx)
}

// GetAccountProfile
//
//	SELECT *
//	FROM profiles
//	LEFT JOIN avatars ON avatars.profile_id = profiles.id
//	WHERE profiles.id = ?
//	LIMIT 1
func (m *Profile) GetAccountProfile(ctx context.Context, profileID uint64) (*ent.Profile, error) {
	return m.Query().
		Where(profile.ID(profileID)).
		WithAvatar().
		Only(ctx)
}

// AccountHiddenCounts
//
//	SELECT *
//	FROM user_settings
//	WHERE user_id = ?
//	LIMIT 1
func (m *Profile) AccountHiddenCounts(ctx context.Context, userID uint64) (hideFollowing bool, hideFollowers bool, err error) {
	settings, err := m.db.UserSetting.Query().
		Where(usersetting.UserID(userID)).
		Only(ctx)
	if ent.IsNotFound(err) {
		return false, false, nil
	}
	if err != nil {
		return false, false, err
	}
	return !settings.ShowProfileFollowingCount, !settings.ShowProfileFollowerCount, nil
}

// IsAdminAccount
//
//	SELECT EXISTS(
//	  SELECT 1 FROM users
//	  WHERE profile_id = ?
//	    AND is_admin = true
//	)
func (m *Profile) IsAdminAccount(ctx context.Context, profileID uint64) (bool, error) {
	return m.db.User.Query().
		Where(
			user.ProfileID(profileID),
			user.IsAdmin(true),
		).
		Exist(ctx)
}

// AccountPronouns
//
//	SELECT *
//	FROM user_pronouns
//	WHERE profile_id = ?
//	LIMIT 1
func (m *Profile) AccountPronouns(ctx context.Context, profileID uint64) (*ent.UserPronoun, error) {
	return m.db.UserPronoun.Query().
		Where(userpronoun.ProfileID(int64(profileID))).
		Only(ctx)
}
