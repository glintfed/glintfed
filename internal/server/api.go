package server

import (
	"glintfed/internal/data"
	"glintfed/internal/server/handler"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/go-chi/chi/v5"
)

func NewAPIServer(cfg *data.Config, handlers *handler.APIHandlers) *http.Server {
	mux := chi.NewRouter()

	// Root Routes
	mux.Post("/f/inbox", handlers.Federation.SharedInbox)
	mux.Post("/users/{username}/inbox", handlers.Federation.UserInbox)
	mux.Get("/i/actor", handlers.InstanceActor.Profile)
	mux.Post("/i/actor/inbox", handlers.InstanceActor.Inbox)
	mux.Get("/i/actor/outbox", handlers.InstanceActor.Outbox)
	mux.Get("/stories/{username}/{id}", handlers.Story.GetActivityObject)

	mux.Get("/.well-known/webfinger", handlers.Federation.Webfinger)
	mux.Get("/.well-known/nodeinfo", handlers.Federation.NodeinfoWellKnown)
	mux.Get("/.well-known/host-meta", handlers.Federation.HostMeta)
	mux.Handle("GET /.well-known/change-password", http.RedirectHandler("/settings/password", http.StatusFound))

	mux.Get("/api/nodeinfo/2.0.json", handlers.Federation.Nodeinfo)
	mux.Get("/api/service/health-check", handlers.HealthCheck.Get)
	mux.Post("/api/auth/app-code-verify", handlers.AppRegister.VerifyCode)
	mux.Post("/api/auth/onboarding", handlers.AppRegister.Onboarding)

	// OAuth2 Routes
	mux.Get("/oauth/authorize", handlers.OAuth.Authorize)
	mux.Post("/oauth/token", handlers.OAuth.Token)
	mux.Post("/oauth/revoke", handlers.OAuth.Revoke)
	mux.Get("/storage/m/_v2/{pid}/{mhash}/{uhash}/{f}", handlers.Media.FallbackRedirect)

	// API Routes
	mux.Route("/api", func(r chi.Router) {

		// V0
		r.Route("/v0/groups", func(r chi.Router) {
			r.Get("/config", handlers.GroupAPI.GetConfig)
			r.Post("/permission/create", handlers.GroupCreate.CheckCreatePermission)
			r.Post("/create", handlers.GroupCreate.StoreGroup)

			r.Post("/search/invite/friends/send", handlers.GroupSearch.InviteFriendsToGroup)
			r.Post("/search/invite/friends", handlers.GroupSearch.SearchFriendsToInvite)
			r.Post("/search/global", handlers.GroupSearch.SearchGlobalResults)
			r.Post("/search/lac", handlers.GroupSearch.SearchLocalAutocomplete)
			r.Post("/search/addrec", handlers.GroupSearch.SearchAddRecent)
			r.Get("/search/getrec", handlers.GroupSearch.SearchGetRecent)

			r.Get("/comments", handlers.GroupComment.GetComments)
			r.Post("/comment", handlers.GroupComment.StoreComment)
			r.Post("/comment/photo", handlers.GroupComment.StoreCommentPhoto)
			r.Post("/comment/delete", handlers.GroupComment.DeleteComment)

			r.Get("/discover/popular", handlers.GroupDiscover.GetDiscoverPopular)
			r.Get("/discover/new", handlers.GroupDiscover.GetDiscoverNew)

			r.Post("/delete", handlers.GroupMeta.DeleteGroup)

			r.Post("/status/new", handlers.GroupPost.StorePost)
			r.Post("/status/delete", handlers.GroupPost.DeletePost)
			r.Post("/status/like", handlers.GroupPost.LikePost)
			r.Post("/status/unlike", handlers.GroupPost.UnlikePost)

			r.Get("/topics/list", handlers.GroupTopic.GroupTopics)
			r.Get("/topics/tag", handlers.GroupTopic.GroupTopicTag)

			r.Get("/accounts/{gid}/{pid}", handlers.GroupAPI.GetGroupAccount)
			r.Get("/categories/list", handlers.GroupAPI.GetGroupCategories)
			r.Get("/category/list", handlers.GroupAPI.GetGroupsByCategory)
			r.Get("/self/recommended/list", handlers.GroupAPI.GetRecommendedGroups)
			r.Get("/self/list", handlers.GroupAPI.GetSelfGroups)

			r.Get("/media/list", handlers.GroupPost.GetGroupMedia)

			r.Get("/members/list", handlers.GroupMember.GetGroupMembers)
			r.Get("/members/requests", handlers.GroupMember.GetGroupMemberJoinRequests)
			r.Post("/members/request", handlers.GroupMember.HandleGroupMemberJoinRequest)
			r.Get("/members/get", handlers.GroupMember.GetGroupMember)
			r.Get("/member/intersect/common", handlers.GroupMember.GetGroupMemberCommonIntersections)

			r.Get("/status", handlers.GroupPost.GetStatus)

			r.Post("/like", handlers.Group.LikePost)
			r.Post("/comment/like", handlers.GroupComment.LikePost)
			r.Post("/comment/unlike", handlers.GroupComment.UnlikePost)

			r.Get("/self/feed", handlers.GroupFeed.GetSelfFeed)
			r.Get("/self/notifications", handlers.GroupNotification.SelfGlobalNotifications)

			r.Get("/{id}/user/{pid}/feed", handlers.GroupFeed.GetGroupProfileFeed)
			r.Get("/{id}/feed", handlers.GroupFeed.GetGroupFeed)
			r.Get("/{id}/atabs", handlers.GroupAdminAPI.GetAdminTabs)
			r.Get("/{id}/admin/interactions", handlers.GroupAdminAPI.GetInteractionLogs)
			r.Get("/{id}/admin/blocks", handlers.GroupAdminAPI.GetBlocks)
			r.Post("/{id}/admin/blocks/add", handlers.GroupAdminAPI.AddBlock)
			r.Post("/{id}/admin/blocks/undo", handlers.GroupAdminAPI.UndoBlock)
			r.Post("/{id}/admin/blocks/export", handlers.GroupAdminAPI.ExportBlocks)
			r.Get("/{id}/reports/list", handlers.GroupAdminAPI.GetReportList)

			r.Get("/{id}/members/interaction-limits", handlers.Group.GetMemberInteractionLimits)
			r.Post("/{id}/invite/check", handlers.Group.GroupMemberInviteCheck)
			r.Post("/{id}/invite/accept", handlers.Group.GroupMemberInviteAccept)
			r.Post("/{id}/invite/decline", handlers.Group.GroupMemberInviteDecline)
			r.Post("/{id}/members/interaction-limits", handlers.Group.UpdateMemberInteractionLimits)
			r.Post("/{id}/report/action", handlers.Group.ReportAction)
			r.Post("/{id}/report/create", handlers.Group.ReportCreate)
			r.Post("/{id}/admin/mbs", handlers.Group.MetaBlockSearch)
			r.Post("/{id}/join", handlers.Group.JoinGroup)
			r.Post("/{id}/cjr", handlers.Group.CancelJoinRequest)
			r.Post("/{id}/leave", handlers.Group.GroupLeave)
			r.Post("/{id}/settings", handlers.Group.UpdateGroup)
			r.Get("/{id}/likes/{sid}", handlers.Group.ShowStatusLikes)
			r.Get("/{id}", handlers.Group.GetGroup)
		})

		// V1
		r.Route("/v1", func(r chi.Router) {
			r.Post("/apps", handlers.APIv1.Apps)
			r.Get("/apps/verify_credentials", handlers.APIv1.GetApp)
			r.Get("/instance", handlers.APIv1.Instance)
			r.Get("/instance/peers", handlers.APIv1.InstancePeers)
			r.Get("/bookmarks", handlers.APIv1.Bookmarks)

			r.Get("/accounts/verify_credentials", handlers.APIv1.VerifyCredentials)
			r.Post("/accounts/update_credentials", handlers.APIv1.AccountUpdateCredentials)
			r.Patch("/accounts/update_credentials", handlers.APIv1.AccountUpdateCredentials)
			r.Get("/accounts/relationships", handlers.APIv1.AccountRelationshipsById)
			r.Get("/accounts/lookup", handlers.APIv1.AccountLookupById)
			r.Get("/accounts/search", handlers.APIv1.AccountSearch)
			r.Get("/accounts/{id}/statuses", handlers.APIv1.AccountStatusesById)
			r.Get("/accounts/{id}/following", handlers.APIv1.AccountFollowingById)
			r.Get("/accounts/{id}/followers", handlers.APIv1.AccountFollowersById)
			r.Post("/accounts/{id}/follow", handlers.APIv1.AccountFollowById)
			r.Post("/accounts/{id}/unfollow", handlers.APIv1.AccountUnfollowById)
			r.Post("/accounts/{id}/block", handlers.APIv1.AccountBlockById)
			r.Post("/accounts/{id}/unblock", handlers.APIv1.AccountUnblockById)
			r.Post("/accounts/{id}/remove_from_followers", handlers.APIv1.AccountRemoveFollowById)
			r.Post("/accounts/{id}/pin", handlers.APIv1.AccountEndorsements)
			r.Post("/accounts/{id}/unpin", handlers.APIv1.AccountEndorsements)
			r.Post("/accounts/{id}/mute", handlers.APIv1.AccountMuteById)
			r.Post("/accounts/{id}/unmute", handlers.APIv1.AccountUnmuteById)
			r.Get("/accounts/{id}/lists", handlers.APIv1.AccountListsById)
			r.Get("/lists/{id}/accounts", handlers.APIv1.AccountListsById)
			r.Get("/accounts/{id}", handlers.APIv1.AccountById)

			r.Post("/avatar/update", handlers.API.AvatarUpdate)
			r.Get("/blocks", handlers.APIv1.AccountBlocks)
			r.Get("/conversations", handlers.APIv1.Conversations)
			r.Get("/custom_emojis", handlers.APIv1.CustomEmojis)
			r.Get("/domain_blocks", handlers.DomainBlock.Index)
			r.Post("/domain_blocks", handlers.DomainBlock.Store)
			r.Delete("/domain_blocks", handlers.DomainBlock.Delete)
			r.Get("/endorsements", handlers.APIv1.AccountEndorsements)
			r.Get("/favourites", handlers.APIv1.AccountFavourites)
			r.Get("/filters", handlers.APIv1.AccountFilters)
			r.Get("/follow_requests", handlers.APIv1.AccountFollowRequests)
			r.Post("/follow_requests/{id}/authorize", handlers.APIv1.AccountFollowRequestAccept)
			r.Post("/follow_requests/{id}/reject", handlers.APIv1.AccountFollowRequestReject)
			r.Get("/lists", handlers.APIv1.AccountLists)
			r.Post("/media", handlers.APIv1.MediaUpload)
			r.Get("/media/{id}", handlers.APIv1.MediaGet)
			r.Put("/media/{id}", handlers.APIv1.MediaUpdate)
			r.Get("/mutes", handlers.APIv1.AccountMutes)
			r.Get("/notifications", handlers.APIv1.AccountNotifications)
			r.Get("/suggestions", handlers.APIv1.AccountSuggestions)

			r.Post("/statuses/{id}/favourite", handlers.APIv1.StatusFavouriteById)
			r.Post("/statuses/{id}/unfavourite", handlers.APIv1.StatusUnfavouriteById)
			r.Get("/statuses/{id}/context", handlers.APIv1.StatusContext)
			r.Get("/statuses/{id}/card", handlers.APIv1.StatusCard)
			r.Get("/statuses/{id}/reblogged_by", handlers.APIv1.StatusRebloggedBy)
			r.Get("/statuses/{id}/favourited_by", handlers.APIv1.StatusFavouritedBy)
			r.Post("/statuses/{id}/reblog", handlers.APIv1.StatusShare)
			r.Post("/statuses/{id}/unreblog", handlers.APIv1.StatusUnshare)
			r.Post("/statuses/{id}/bookmark", handlers.APIv1.BookmarkStatus)
			r.Post("/statuses/{id}/unbookmark", handlers.APIv1.UnbookmarkStatus)
			r.Post("/statuses/{id}/pin", handlers.APIv1.StatusPin)
			r.Post("/statuses/{id}/unpin", handlers.APIv1.StatusUnpin)
			r.Delete("/statuses/{id}", handlers.APIv1.StatusDelete)
			r.Get("/statuses/{id}", handlers.APIv1.StatusById)
			r.Post("/statuses", handlers.APIv1.StatusCreate)

			r.Get("/timelines/home", handlers.APIv1.TimelineHome)
			r.Get("/timelines/public", handlers.APIv1.TimelinePublic)
			r.Get("/timelines/tag/{hashtag}", handlers.APIv1.TimelineHashtag)
			r.Get("/discover/posts", handlers.APIv1.DiscoverPosts)

			r.Get("/preferences", handlers.APIv1.GetPreferences)
			r.Get("/trends", handlers.APIv1.GetTrends)
			r.Get("/announcements", handlers.APIv1.GetAnnouncements)
			r.Get("/markers", handlers.APIv1.GetMarkers)
			r.Post("/markers", handlers.APIv1.SetMarkers)

			r.Get("/followed_tags", handlers.Tags.GetFollowedTags)
			r.Post("/tags/{id}/follow", handlers.Tags.FollowHashtag)
			r.Post("/tags/{id}/unfollow", handlers.Tags.UnfollowHashtag)
			r.Get("/tags/{id}/related", handlers.Tags.RelatedTags)
			r.Get("/tags/{id}", handlers.Tags.GetHashtag)

			r.Get("/statuses/{id}/history", handlers.StatusEdit.History)
			r.Put("/statuses/{id}", handlers.StatusEdit.Store)

			r.Route("/admin", func(r chi.Router) {
				r.Get("/domain_blocks", handlers.AdminDomainBlock.Index)
				r.Post("/domain_blocks", handlers.AdminDomainBlock.Create)
				r.Get("/domain_blocks/{id}", handlers.AdminDomainBlock.Show)
				r.Put("/domain_blocks/{id}", handlers.AdminDomainBlock.Update)
				r.Delete("/domain_blocks/{id}", handlers.AdminDomainBlock.Delete)
			})
		})

		// V2
		r.Route("/v2", func(r chi.Router) {
			r.Get("/search", handlers.APIv2.Search)
			r.Post("/media", handlers.APIv2.MediaUploadV2)
			r.Get("/streaming/config", handlers.APIv2.GetWebsocketConfig)
			r.Get("/instance", handlers.APIv2.Instance)

			r.Get("/filters", handlers.CustomFilter.Index)
			r.Get("/filters/{id}", handlers.CustomFilter.Show)
			r.Post("/filters", handlers.CustomFilter.Store)
			r.Put("/filters/{id}", handlers.CustomFilter.Update)
			r.Delete("/filters/{id}", handlers.CustomFilter.Delete)
		})

		// V1.1
		r.Route("/v1.1", func(r chi.Router) {
			r.Post("/report", handlers.APIv1Dot1.Report)

			r.Route("/accounts", func(r chi.Router) {
				r.Get("/timelines/home", handlers.APIv1.TimelineHome)
				r.Delete("/avatar", handlers.APIv1Dot1.DeleteAvatar)
				r.Get("/{id}/posts", handlers.APIv1Dot1.AccountPosts)
				r.Post("/change-password", handlers.APIv1Dot1.AccountChangePassword)
				r.Get("/login-activity", handlers.APIv1Dot1.AccountLoginActivity)
				r.Get("/two-factor", handlers.APIv1Dot1.AccountTwoFactor)
				r.Get("/emails-from-pixelfed", handlers.APIv1Dot1.AccountEmailsFromPixelfed)
				r.Get("/apps-and-applications", handlers.APIv1Dot1.AccountApps)
				r.Get("/mutuals/{id}", handlers.APIv1Dot1.GetMutualAccounts)
				r.Get("/username/{username}", handlers.APIv1Dot1.AccountUsernameToId)
			})

			r.Route("/collections", func(r chi.Router) {
				r.Get("/accounts/{id}", handlers.Collection.GetUserCollections)
				r.Get("/items/{id}", handlers.Collection.GetItems)
				r.Get("/view/{id}", handlers.Collection.GetCollection)
				r.Post("/add", handlers.Collection.StoreId)
				r.Post("/update/{id}", handlers.Collection.Store)
				r.Delete("/delete/{id}", handlers.Collection.Delete)
				r.Post("/remove", handlers.Collection.DeleteId)
				r.Get("/self", handlers.Collection.GetSelfCollections)
			})

			r.Route("/direct", func(r chi.Router) {
				r.Get("/thread", handlers.DirectMessage.Thread)
				r.Post("/thread/send", handlers.DirectMessage.Create)
				r.Delete("/thread/message", handlers.DirectMessage.Delete)
				r.Post("/thread/mute", handlers.DirectMessage.Mute)
				r.Post("/thread/unmute", handlers.DirectMessage.Unmute)
				r.Post("/thread/media", handlers.DirectMessage.MediaUpload)
				r.Post("/thread/read", handlers.DirectMessage.Read)
				r.Post("/lookup", handlers.DirectMessage.ComposeLookup)
				r.Get("/compose/mutuals", handlers.DirectMessage.ComposeMutuals)
			})

			r.Route("/archive", func(r chi.Router) {
				r.Post("/add/{id}", handlers.APIv1Dot1.Archive)
				r.Post("/remove/{id}", handlers.APIv1Dot1.Unarchive)
				r.Get("/list", handlers.APIv1Dot1.ArchivedPosts)
			})

			r.Route("/places", func(r chi.Router) {
				r.Get("/posts/{id}/{slug}", handlers.APIv1Dot1.PlacesById)
			})

			r.Route("/stories", func(r chi.Router) {
				r.Get("/carousel", handlers.StoryAPIv1.Carousel)
				r.Post("/add", handlers.StoryAPIv1.Add)
				r.Post("/publish", handlers.StoryAPIv1.Publish)
				r.Post("/seen", handlers.StoryAPIv1.Viewed)
				r.Post("/self-expire/{id}", handlers.StoryAPIv1.Delete)
				r.Post("/comment", handlers.StoryAPIv1.Comment)
			})

			r.Route("/compose", func(r chi.Router) {
				r.Get("/search/location", handlers.Compose.SearchLocation)
				r.Get("/settings", handlers.Compose.ComposeSettings)
			})

			r.Route("/discover", func(r chi.Router) {
				r.Get("/accounts/popular", handlers.APIv1.DiscoverAccountsPopular)
				r.Get("/posts/trending", handlers.Discover.TrendingApi)
				r.Get("/posts/hashtags", handlers.Discover.TrendingHashtags)
				r.Get("/posts/network/trending", handlers.Discover.DiscoverNetworkTrending)
			})

			r.Route("/directory", func(r chi.Router) {
				r.Get("/listing", handlers.PixelfedDirectory.Get)
			})

			r.Route("/auth", func(r chi.Router) {
				r.Get("/iarpfc", handlers.APIv1Dot1.InAppRegistrationPreFlightCheck)
				r.Post("/iar", handlers.APIv1Dot1.InAppRegistration)
				r.Post("/iarc", handlers.APIv1Dot1.InAppRegistrationConfirm)
				r.Get("/iarer", handlers.APIv1Dot1.InAppRegistrationEmailRedirect)

				r.Post("/invite/admin/verify", handlers.AdminInvite.ApiVerifyCheck)
				r.Post("/invite/admin/uc", handlers.AdminInvite.ApiUsernameCheck)
				r.Post("/invite/admin/ec", handlers.AdminInvite.ApiEmailCheck)
			})

			r.Route("/push", func(r chi.Router) {
				r.Get("/state", handlers.APIv1Dot1.GetPushState)
				r.Post("/compare", handlers.APIv1Dot1.ComparePush)
				r.Post("/update", handlers.APIv1Dot1.UpdatePush)
				r.Post("/disable", handlers.APIv1Dot1.DisablePush)
			})

			r.Post("/status/create", handlers.APIv1Dot1.StatusCreate)
			r.Get("/nag/state", handlers.APIv1Dot1.NagState)
		})

		// V1.2
		r.Route("/v1.2", func(r chi.Router) {
			r.Route("/stories", func(r chi.Router) {
				r.Get("/viewers", handlers.StoryAPIv1.Viewers)
				r.Post("/publish", handlers.StoryAPIv1.PublishNext)
				r.Get("/carousel", handlers.StoryAPIv1.CarouselNext)
				r.Get("/mention-autocomplete", handlers.StoryAPIv1.MentionAutocomplete)
			})
		})

		// Admin
		r.Route("/admin", func(r chi.Router) {
			r.Post("/moderate/post/{id}", handlers.APIv1Dot1.ModeratePost)
			r.Get("/supported", handlers.AdminAPI.Supported)
			r.Get("/stats", handlers.AdminAPI.GetStats)

			r.Get("/autospam/list", handlers.AdminAPI.Autospam)
			r.Post("/autospam/handle", handlers.AdminAPI.AutospamHandle)
			r.Get("/mod-reports/list", handlers.AdminAPI.ModReports)
			r.Post("/mod-reports/handle", handlers.AdminAPI.ModReportHandle)
			r.Get("/config", handlers.AdminAPI.GetConfiguration)
			r.Post("/config/update", handlers.AdminAPI.UpdateConfiguration)
			r.Get("/users/list", handlers.AdminAPI.GetUsers)
			r.Get("/users/get", handlers.AdminAPI.GetUser)
			r.Post("/users/action", handlers.AdminAPI.UserAdminAction)
			r.Get("/instances/list", handlers.AdminAPI.Instances)
			r.Get("/instances/get", handlers.AdminAPI.GetInstance)
			r.Post("/instances/moderate", handlers.AdminAPI.ModerateInstance)
			r.Post("/instances/refresh-stats", handlers.AdminAPI.RefreshInstanceStats)
			r.Get("/instance/stats", handlers.AdminAPI.GetAllStats)
		})

		// Landing
		r.Route("/landing/v1", func(r chi.Router) {
			r.Get("/directory", handlers.Landing.GetDirectoryApi)
		})

		// Pixelfed
		r.Route("/pixelfed", func(r chi.Router) {
			r.Route("/v1", func(r chi.Router) {
				r.Post("/report", handlers.APIv1Dot1.Report)

				r.Route("/accounts", func(r chi.Router) {
					r.Get("/timelines/home", handlers.APIv1.TimelineHome)
					r.Delete("/avatar", handlers.APIv1Dot1.DeleteAvatar)
					r.Get("/{id}/posts", handlers.APIv1Dot1.AccountPosts)
					r.Post("/change-password", handlers.APIv1Dot1.AccountChangePassword)
					r.Get("/login-activity", handlers.APIv1Dot1.AccountLoginActivity)
					r.Get("/two-factor", handlers.APIv1Dot1.AccountTwoFactor)
					r.Get("/emails-from-pixelfed", handlers.APIv1Dot1.AccountEmailsFromPixelfed)
					r.Get("/apps-and-applications", handlers.APIv1Dot1.AccountApps)
				})

				r.Route("/archive", func(r chi.Router) {
					r.Post("/add/{id}", handlers.APIv1Dot1.Archive)
					r.Post("/remove/{id}", handlers.APIv1Dot1.Unarchive)
					r.Get("/list", handlers.APIv1Dot1.ArchivedPosts)
				})

				r.Route("/collections", func(r chi.Router) {
					r.Get("/accounts/{id}", handlers.Collection.GetUserCollections)
					r.Get("/items/{id}", handlers.Collection.GetItems)
					r.Get("/view/{id}", handlers.Collection.GetCollection)
					r.Post("/add", handlers.Collection.StoreId)
					r.Post("/update/{id}", handlers.Collection.Store)
					r.Delete("/delete/{id}", handlers.Collection.Delete)
					r.Post("/remove", handlers.Collection.DeleteId)
					r.Get("/self", handlers.Collection.GetSelfCollections)
				})

				r.Route("/compose", func(r chi.Router) {
					r.Get("/search/location", handlers.Compose.SearchLocation)
					r.Get("/settings", handlers.Compose.ComposeSettings)
				})

				r.Route("/direct", func(r chi.Router) {
					r.Get("/thread", handlers.DirectMessage.Thread)
					r.Post("/thread/send", handlers.DirectMessage.Create)
					r.Delete("/thread/message", handlers.DirectMessage.Delete)
					r.Post("/thread/mute", handlers.DirectMessage.Mute)
					r.Post("/thread/unmute", handlers.DirectMessage.Unmute)
					r.Post("/thread/media", handlers.DirectMessage.MediaUpload)
					r.Post("/thread/read", handlers.DirectMessage.Read)
					r.Post("/lookup", handlers.DirectMessage.ComposeLookup)
				})

				r.Route("/discover", func(r chi.Router) {
					r.Get("/accounts/popular", handlers.APIv1.DiscoverAccountsPopular)
					r.Get("/posts/trending", handlers.Discover.TrendingApi)
					r.Get("/posts/hashtags", handlers.Discover.TrendingHashtags)
				})

				r.Route("/directory", func(r chi.Router) {
					r.Get("/listing", handlers.PixelfedDirectory.Get)
				})

				r.Route("/places", func(r chi.Router) {
					r.Get("/posts/{id}/{slug}", handlers.APIv1Dot1.PlacesById)
				})

				r.Get("/web/settings", handlers.APIv1Dot1.GetWebSettings)
				r.Post("/web/settings", handlers.APIv1Dot1.SetWebSettings)
				r.Get("/app/settings", handlers.UserAppSetting.Get)
				r.Post("/app/settings", handlers.UserAppSetting.Store)

				r.Route("/stories", func(r chi.Router) {
					r.Get("/carousel", handlers.StoryAPIv1.Carousel)
					r.Get("/self-carousel", handlers.StoryAPIv1.SelfCarousel)
					r.Post("/add", handlers.StoryAPIv1.Add)
					r.Post("/publish", handlers.StoryAPIv1.Publish)
					r.Post("/seen", handlers.StoryAPIv1.Viewed)
					r.Post("/self-expire/{id}", handlers.StoryAPIv1.Delete)
					r.Post("/comment", handlers.StoryAPIv1.Comment)
					r.Get("/viewers", handlers.StoryAPIv1.Viewers)
				})
			})
		})
	})

	return &http.Server{
		Addr:    cfg.Server.API.Addr,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
}
