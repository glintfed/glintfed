package handler

import (
	"glintfed/internal/server/handler/admininvite"
	"glintfed/internal/server/handler/api"
	"glintfed/internal/server/handler/api/adminapi"
	"glintfed/internal/server/handler/api/apiv1"
	"glintfed/internal/server/handler/api/apiv1/domainblock"
	"glintfed/internal/server/handler/api/apiv1/tags"
	"glintfed/internal/server/handler/api/apiv1dot1"
	"glintfed/internal/server/handler/api/apiv2"
	"glintfed/internal/server/handler/appregister"
	"glintfed/internal/server/handler/collection"
	"glintfed/internal/server/handler/compose"
	"glintfed/internal/server/handler/customfilter"
	"glintfed/internal/server/handler/directmessage"
	"glintfed/internal/server/handler/discover"
	"glintfed/internal/server/handler/federation"
	"glintfed/internal/server/handler/group"
	"glintfed/internal/server/handler/healthcheck"
	"glintfed/internal/server/handler/instanceactor"
	"glintfed/internal/server/handler/landing"
	"glintfed/internal/server/handler/media"
	"glintfed/internal/server/handler/oauth"
	"glintfed/internal/server/handler/pixelfeddirectory"
	"glintfed/internal/server/handler/statusedit"
	"glintfed/internal/server/handler/stories/storyapiv1"
	"glintfed/internal/server/handler/story"
	"glintfed/internal/server/handler/userappsettings"

	admindomainblocks "glintfed/internal/server/handler/api/apiv1/admin/domainblocks"
	groupsadminapi "glintfed/internal/server/handler/groups/admin"
	groupsapi "glintfed/internal/server/handler/groups/api"
	groupscomment "glintfed/internal/server/handler/groups/comment"
	groupscreate "glintfed/internal/server/handler/groups/create"
	groupsdiscover "glintfed/internal/server/handler/groups/discover"
	groupsfeed "glintfed/internal/server/handler/groups/feed"
	groupsmember "glintfed/internal/server/handler/groups/member"
	groupsmeta "glintfed/internal/server/handler/groups/meta"
	groupsnotifications "glintfed/internal/server/handler/groups/notifications"
	groupspost "glintfed/internal/server/handler/groups/post"
	groupssearch "glintfed/internal/server/handler/groups/search"
	groupstopic "glintfed/internal/server/handler/groups/topic"
)

type APIHandlers struct {
	HealthCheck       healthcheck.Handler
	OAuth             oauth.Handler
	Federation        federation.Handler
	InstanceActor     instanceactor.Handler
	Story             story.Handler
	Media             media.Handler
	AppRegister       appregister.Handler
	API               api.Handler
	APIv1             apiv1.Handler
	APIv1Dot1         apiv1dot1.Handler
	APIv2             apiv2.Handler
	Tags              tags.Handler
	DomainBlock       domainblock.Handler
	StatusEdit        statusedit.Handler
	AdminDomainBlock  admindomainblocks.Handler
	CustomFilter      customfilter.Handler
	Discover          discover.Handler
	PixelfedDirectory pixelfeddirectory.Handler
	StoryAPIv1        storyapiv1.Handler
	Compose           compose.Handler
	Landing           landing.Handler
	AdminInvite       admininvite.Handler
	UserAppSetting    userappsettings.Handler
	AdminAPI          adminapi.Handler
	Collection        collection.Handler
	DirectMessage     directmessage.Handler
	GroupAPI          groupsapi.Handler
	GroupCreate       groupscreate.Handler
	GroupSearch       groupssearch.Handler
	GroupComment      groupscomment.Handler
	GroupDiscover     groupsdiscover.Handler
	GroupMeta         groupsmeta.Handler
	GroupPost         groupspost.Handler
	GroupTopic        groupstopic.Handler
	GroupMember       groupsmember.Handler
	GroupFeed         groupsfeed.Handler
	GroupNotification groupsnotifications.Handler
	GroupAdminAPI     groupsadminapi.Handler
	Group             group.Handler
}

func NewAPIHandlers(
	healthCheck healthcheck.Handler,
	oauthSvc oauth.Handler,
	federation federation.Handler,
	instanceActor instanceactor.Handler,
	story story.Handler,
	media media.Handler,
	appRegister appregister.Handler,
	api api.Handler,
	apiv1 apiv1.Handler,
	apiv1dot1 apiv1dot1.Handler,
	apiv2 apiv2.Handler,
	tags tags.Handler,
	domainBlock domainblock.Handler,
	statusEdit statusedit.Handler,
	adminDomainBlock admindomainblocks.Handler,
	customFilter customfilter.Handler,
	discover discover.Handler,
	pixelfedDirectory pixelfeddirectory.Handler,
	storyAPIv1 storyapiv1.Handler,
	compose compose.Handler,
	landing landing.Handler,
	adminInvite admininvite.Handler,
	userAppSetting userappsettings.Handler,
	adminAPI adminapi.Handler,
	collection collection.Handler,
	directMessage directmessage.Handler,
	groupAPI groupsapi.Handler,
	groupCreate groupscreate.Handler,
	groupSearch groupssearch.Handler,
	groupComment groupscomment.Handler,
	groupDiscover groupsdiscover.Handler,
	groupMeta groupsmeta.Handler,
	groupPost groupspost.Handler,
	groupTopic groupstopic.Handler,
	groupMember groupsmember.Handler,
	groupFeed groupsfeed.Handler,
	groupNotification groupsnotifications.Handler,
	groupAdminAPI groupsadminapi.Handler,
	group group.Handler,
) *APIHandlers {
	return &APIHandlers{
		HealthCheck:       healthCheck,
		OAuth:             oauthSvc,
		Federation:        federation,
		InstanceActor:     instanceActor,
		Story:             story,
		Media:             media,
		AppRegister:       appRegister,
		API:               api,
		APIv1:             apiv1,
		APIv1Dot1:         apiv1dot1,
		APIv2:             apiv2,
		Tags:              tags,
		DomainBlock:       domainBlock,
		StatusEdit:        statusEdit,
		AdminDomainBlock:  adminDomainBlock,
		CustomFilter:      customFilter,
		Discover:          discover,
		PixelfedDirectory: pixelfedDirectory,
		StoryAPIv1:        storyAPIv1,
		Compose:           compose,
		Landing:           landing,
		AdminInvite:       adminInvite,
		UserAppSetting:    userAppSetting,
		AdminAPI:          adminAPI,
		Collection:        collection,
		DirectMessage:     directMessage,
		GroupAPI:          groupAPI,
		GroupCreate:       groupCreate,
		GroupSearch:       groupSearch,
		GroupComment:      groupComment,
		GroupDiscover:     groupDiscover,
		GroupMeta:         groupMeta,
		GroupPost:         groupPost,
		GroupTopic:        groupTopic,
		GroupMember:       groupMember,
		GroupFeed:         groupFeed,
		GroupNotification: groupNotification,
		GroupAdminAPI:     groupAdminAPI,
		Group:             group,
	}
}
