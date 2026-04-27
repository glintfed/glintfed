package apiv1

import (
	"net/http"

	"glintfed/internal/server/handler/internal"
)

type Handler interface {
	Apps(w http.ResponseWriter, r *http.Request)
	GetApp(w http.ResponseWriter, r *http.Request)
	Instance(w http.ResponseWriter, r *http.Request)
	InstancePeers(w http.ResponseWriter, r *http.Request)
	Bookmarks(w http.ResponseWriter, r *http.Request)
	VerifyCredentials(w http.ResponseWriter, r *http.Request)
	AccountUpdateCredentials(w http.ResponseWriter, r *http.Request)
	AccountRelationshipsById(w http.ResponseWriter, r *http.Request)
	AccountLookupById(w http.ResponseWriter, r *http.Request)
	AccountSearch(w http.ResponseWriter, r *http.Request)
	AccountStatusesById(w http.ResponseWriter, r *http.Request)
	AccountFollowingById(w http.ResponseWriter, r *http.Request)
	AccountFollowersById(w http.ResponseWriter, r *http.Request)
	AccountFollowById(w http.ResponseWriter, r *http.Request)
	AccountUnfollowById(w http.ResponseWriter, r *http.Request)
	AccountBlockById(w http.ResponseWriter, r *http.Request)
	AccountUnblockById(w http.ResponseWriter, r *http.Request)
	AccountRemoveFollowById(w http.ResponseWriter, r *http.Request)
	AccountEndorsements(w http.ResponseWriter, r *http.Request)
	AccountMuteById(w http.ResponseWriter, r *http.Request)
	AccountUnmuteById(w http.ResponseWriter, r *http.Request)
	AccountListsById(w http.ResponseWriter, r *http.Request)
	AccountById(w http.ResponseWriter, r *http.Request)
	AccountBlocks(w http.ResponseWriter, r *http.Request)
	Conversations(w http.ResponseWriter, r *http.Request)
	CustomEmojis(w http.ResponseWriter, r *http.Request)
	AccountFavourites(w http.ResponseWriter, r *http.Request)
	AccountFilters(w http.ResponseWriter, r *http.Request)
	AccountFollowRequests(w http.ResponseWriter, r *http.Request)
	AccountFollowRequestAccept(w http.ResponseWriter, r *http.Request)
	AccountFollowRequestReject(w http.ResponseWriter, r *http.Request)
	AccountLists(w http.ResponseWriter, r *http.Request)
	MediaUpload(w http.ResponseWriter, r *http.Request)
	MediaGet(w http.ResponseWriter, r *http.Request)
	MediaUpdate(w http.ResponseWriter, r *http.Request)
	AccountMutes(w http.ResponseWriter, r *http.Request)
	AccountNotifications(w http.ResponseWriter, r *http.Request)
	AccountSuggestions(w http.ResponseWriter, r *http.Request)
	StatusFavouriteById(w http.ResponseWriter, r *http.Request)
	StatusUnfavouriteById(w http.ResponseWriter, r *http.Request)
	StatusContext(w http.ResponseWriter, r *http.Request)
	StatusCard(w http.ResponseWriter, r *http.Request)
	StatusRebloggedBy(w http.ResponseWriter, r *http.Request)
	StatusFavouritedBy(w http.ResponseWriter, r *http.Request)
	StatusShare(w http.ResponseWriter, r *http.Request)
	StatusUnshare(w http.ResponseWriter, r *http.Request)
	BookmarkStatus(w http.ResponseWriter, r *http.Request)
	UnbookmarkStatus(w http.ResponseWriter, r *http.Request)
	StatusPin(w http.ResponseWriter, r *http.Request)
	StatusUnpin(w http.ResponseWriter, r *http.Request)
	StatusDelete(w http.ResponseWriter, r *http.Request)
	StatusById(w http.ResponseWriter, r *http.Request)
	StatusCreate(w http.ResponseWriter, r *http.Request)
	TimelineHome(w http.ResponseWriter, r *http.Request)
	TimelinePublic(w http.ResponseWriter, r *http.Request)
	TimelineHashtag(w http.ResponseWriter, r *http.Request)
	DiscoverPosts(w http.ResponseWriter, r *http.Request)
	GetPreferences(w http.ResponseWriter, r *http.Request)
	GetTrends(w http.ResponseWriter, r *http.Request)
	GetAnnouncements(w http.ResponseWriter, r *http.Request)
	GetMarkers(w http.ResponseWriter, r *http.Request)
	SetMarkers(w http.ResponseWriter, r *http.Request)
	DiscoverAccountsPopular(w http.ResponseWriter, r *http.Request)
}

func New() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Apps(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Apps")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetApp(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetApp")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Instance(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Instance")
	defer span.End()
	// TODO: Implement
}

func (h *handler) InstancePeers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.InstancePeers")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Bookmarks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Bookmarks")
	defer span.End()
	// TODO: Implement
}

func (h *handler) VerifyCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.VerifyCredentials")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountUpdateCredentials(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUpdateCredentials")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountRelationshipsById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountRelationshipsById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountLookupById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountLookupById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountSearch(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountSearch")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountStatusesById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountStatusesById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowingById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowingById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowersById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowersById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountUnfollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnfollowById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountBlockById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountBlockById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountUnblockById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnblockById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountRemoveFollowById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountRemoveFollowById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountEndorsements(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountEndorsements")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountMuteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountMutesById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountUnmuteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountUnmutedById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountListsById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountListsById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountBlocks(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountBlocks")
	defer span.End()
	// TODO: Implement
}

func (h *handler) Conversations(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.Conversations")
	defer span.End()
	// TODO: Implement
}

func (h *handler) CustomEmojis(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.CustomEmojis")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFavourites(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFavourites")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFilters(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFilters")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowRequests(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequest")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowRequestAccept(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequestAccept")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountFollowRequestReject(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountFollowRequestReject")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountLists(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountLists")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MediaUpload(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaUpload")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MediaGet(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaGet")
	defer span.End()
	// TODO: Implement
}

func (h *handler) MediaUpdate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.MediaUpdate")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountMutes(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountMutes")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountNotifications(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountNotifications")
	defer span.End()
	// TODO: Implement
}

func (h *handler) AccountSuggestions(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.AccountSuggestions")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusFavouriteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusFavouriteById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusUnfavouriteById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnfavouriteById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusContext(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusContext")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusCard(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusCard")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusRebloggedBy(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusRebloggedBy")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusFavouritedBy(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusFavouritedBy")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusShare(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusShare")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusUnshare(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnshare")
	defer span.End()
	// TODO: Implement
}

func (h *handler) BookmarkStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.BookmarkStatus")
	defer span.End()
	// TODO: Implement
}

func (h *handler) UnbookmarkStatus(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.UnbookmarkStatus")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusPin(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusPin")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusUnpin(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusUnpin")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusDelete(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusDelete")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusById(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusById")
	defer span.End()
	// TODO: Implement
}

func (h *handler) StatusCreate(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.StatusCreate")
	defer span.End()
	// TODO: Implement
}

func (h *handler) TimelineHome(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelineHome")
	defer span.End()
	// TODO: Implement
}

func (h *handler) TimelinePublic(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelinePublic")
	defer span.End()
	// TODO: Implement
}

func (h *handler) TimelineHashtag(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.TimelineHashtag")
	defer span.End()
	// TODO: Implement
}

func (h *handler) DiscoverPosts(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DiscoverPosts")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetPreferences(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetPreferences")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetTrends(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetTrends")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetAnnouncements(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetAnnouncements")
	defer span.End()
	// TODO: Implement
}

func (h *handler) GetMarkers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.GetMarkers")
	defer span.End()
	// TODO: Implement
}

func (h *handler) SetMarkers(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.SetMarkers")
	defer span.End()
	// TODO: Implement
}

func (h *handler) DiscoverAccountsPopular(w http.ResponseWriter, r *http.Request) {
	_, span := internal.T.Start(r.Context(), "Api.ApiV1.DiscoverAccountsPopular")
	defer span.End()
	// TODO: Implement
}
