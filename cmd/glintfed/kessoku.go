package main

import (
	"glintfed/internal/data/client"
	"glintfed/internal/lib/fositestore"
	"glintfed/internal/model"
	oauthrepo "glintfed/internal/repo/oauth"
	"glintfed/internal/server"
	"glintfed/internal/server/handler"
	"glintfed/internal/server/handler/admininvite"
	apiroot "glintfed/internal/server/handler/api"
	"glintfed/internal/server/handler/api/adminapi"
	"glintfed/internal/server/handler/api/apiv1"
	apiv1admindomainblocks "glintfed/internal/server/handler/api/apiv1/admin/domainblocks"
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
	"glintfed/internal/server/handler/groups/admin"
	groupsapi "glintfed/internal/server/handler/groups/api"
	"glintfed/internal/server/handler/groups/comment"
	groupscreate "glintfed/internal/server/handler/groups/create"
	groupsdiscover "glintfed/internal/server/handler/groups/discover"
	"glintfed/internal/server/handler/groups/feed"
	"glintfed/internal/server/handler/groups/member"
	"glintfed/internal/server/handler/groups/meta"
	"glintfed/internal/server/handler/groups/notifications"
	"glintfed/internal/server/handler/groups/post"
	groupssearch "glintfed/internal/server/handler/groups/search"
	"glintfed/internal/server/handler/groups/topic"
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
	accountsvc "glintfed/internal/service/account"
	instancesvc "glintfed/internal/service/instance"
	"glintfed/internal/service/worker"
	"net/http"

	"github.com/mazrean/kessoku"
)

//go:generate go tool kessoku $GOFILE
var _ = kessoku.Inject[*App](
	"newApp",
	kessoku.Set(
		kessoku.Provide(client.NewDatabase),
		kessoku.Provide(client.NewEvent),
		kessoku.Provide(server.NewAPIServer),
	),
	kessoku.Set(
		kessoku.Provide(fositestore.New),
		kessoku.Provide(fositestore.NewOAuth2Provider),
	),
	kessoku.Set(
		kessoku.Provide(model.NewAppRegister),
		kessoku.Provide(model.NewUser),
		kessoku.Provide(model.NewInstanceActor),
		kessoku.Provide(model.NewInstance),
		kessoku.Provide(model.NewProfile),
		kessoku.Provide(model.NewStatus),
		kessoku.Provide(model.NewMedia),
	),
	kessoku.Set(
		kessoku.Bind[instancesvc.InstanceModel](
			kessoku.Provide(func(m *model.Instance) instancesvc.InstanceModel { return m }),
		),
		kessoku.Bind[instancesvc.NodeinfoUserModel](
			kessoku.Provide(func(m *model.User) instancesvc.NodeinfoUserModel { return m }),
		),
		kessoku.Bind[instancesvc.NodeinfoStatusModel](
			kessoku.Provide(func(m *model.Status) instancesvc.NodeinfoStatusModel { return m }),
		),
		kessoku.Provide(instancesvc.NewDomainManager),
		kessoku.Provide(instancesvc.NewNodeInfo),
	),
	kessoku.Set(
		kessoku.Bind[accountsvc.ProfileModel](
			kessoku.Provide(func(m *model.Profile) accountsvc.ProfileModel { return m }),
		),
		kessoku.Provide(accountsvc.NewProfile),
	),
	kessoku.Set(
		kessoku.Bind[appregister.AppRegisterModel](
			kessoku.Provide(func(m *model.AppRegister) appregister.AppRegisterModel { return m }),
		),
		kessoku.Bind[appregister.UserModel](
			kessoku.Provide(func(m *model.User) appregister.UserModel { return m }),
		),
		kessoku.Provide(oauthrepo.NewRefreshToken),
		kessoku.Bind[appregister.RefreshTokenRepository](
			kessoku.Provide(func(r *oauthrepo.RefreshToken) appregister.RefreshTokenRepository { return r }),
		),
		kessoku.Bind[appregister.AccountService](
			kessoku.Provide(func(s *accountsvc.Profile) appregister.AccountService { return s }),
		),
	),
	kessoku.Set(
		kessoku.Bind[instanceactor.InstanceActorModel](
			kessoku.Provide(func(m *model.InstanceActor) instanceactor.InstanceActorModel { return m }),
		),
	),
	kessoku.Set(
		kessoku.Bind[story.ProfileModel](
			kessoku.Provide(func(m *model.Profile) story.ProfileModel { return m }),
		),
	),
	kessoku.Set(
		kessoku.Bind[media.MediaModel](
			kessoku.Provide(func(m *model.Media) media.MediaModel { return m }),
		),
	),
	kessoku.Set(
		kessoku.Bind[federation.ProfileModel](
			kessoku.Provide(func(m *model.Profile) federation.ProfileModel { return m }),
		),
		kessoku.Bind[federation.StatusModel](
			kessoku.Provide(func(m *model.Status) federation.StatusModel { return m }),
		),
		kessoku.Provide(worker.NewInboxWorker),
		kessoku.Bind[federation.InboxWorkerService](
			kessoku.Provide(func(s *worker.InboxWorker) federation.InboxWorkerService { return s }),
		),
		kessoku.Bind[federation.InstanceService](
			kessoku.Provide(func(dm *instancesvc.DomainManager, n *instancesvc.Nodeinfo) federation.InstanceService {
				return &struct {
					*instancesvc.DomainManager
					*instancesvc.Nodeinfo
				}{
					DomainManager: dm,
					Nodeinfo:      n,
				}
			}),
		),
	),
	kessoku.Set(
		kessoku.Bind[oauth.UserModel](
			kessoku.Provide(func(m *model.User) oauth.UserModel { return m }),
		),
	),
	kessoku.Set(
		kessoku.Bind[admininvite.Handler](kessoku.Provide(admininvite.New)),
		kessoku.Bind[apiroot.Handler](kessoku.Provide(apiroot.New)),
		kessoku.Bind[adminapi.Handler](kessoku.Provide(adminapi.New)),
		kessoku.Bind[apiv1.Handler](kessoku.Provide(apiv1.New)),
		kessoku.Bind[apiv1admindomainblocks.Handler](kessoku.Provide(apiv1admindomainblocks.New)),
		kessoku.Bind[domainblock.Handler](kessoku.Provide(domainblock.New)),
		kessoku.Bind[tags.Handler](kessoku.Provide(tags.New)),
		kessoku.Bind[apiv1dot1.Handler](kessoku.Provide(apiv1dot1.New)),
		kessoku.Bind[apiv2.Handler](kessoku.Provide(apiv2.New)),
		kessoku.Bind[appregister.Handler](kessoku.Provide(appregister.New)),
		kessoku.Bind[collection.Handler](kessoku.Provide(collection.New)),
		kessoku.Bind[compose.Handler](kessoku.Provide(compose.New)),
		kessoku.Bind[customfilter.Handler](kessoku.Provide(customfilter.New)),
		kessoku.Bind[directmessage.Handler](kessoku.Provide(directmessage.New)),
		kessoku.Bind[discover.Handler](kessoku.Provide(discover.New)),
		kessoku.Bind[federation.Handler](kessoku.Provide(federation.New)),
		kessoku.Bind[group.Handler](kessoku.Provide(group.New)),
		kessoku.Bind[admin.Handler](kessoku.Provide(admin.New)),
		kessoku.Bind[groupsapi.Handler](kessoku.Provide(groupsapi.New)),
		kessoku.Bind[comment.Handler](kessoku.Provide(comment.New)),
		kessoku.Bind[groupscreate.Handler](kessoku.Provide(groupscreate.New)),
		kessoku.Bind[groupsdiscover.Handler](kessoku.Provide(groupsdiscover.New)),
		kessoku.Bind[feed.Handler](kessoku.Provide(feed.New)),
		kessoku.Bind[member.Handler](kessoku.Provide(member.New)),
		kessoku.Bind[meta.Handler](kessoku.Provide(meta.New)),
		kessoku.Bind[notifications.Handler](kessoku.Provide(notifications.New)),
		kessoku.Bind[post.Handler](kessoku.Provide(post.New)),
		kessoku.Bind[groupssearch.Handler](kessoku.Provide(groupssearch.New)),
		kessoku.Bind[topic.Handler](kessoku.Provide(topic.New)),
		kessoku.Bind[healthcheck.Handler](kessoku.Provide(healthcheck.New)),
		kessoku.Bind[instanceactor.Handler](kessoku.Provide(instanceactor.New)),
		kessoku.Bind[landing.Handler](kessoku.Provide(landing.New)),
		kessoku.Bind[media.Handler](kessoku.Provide(media.New)),
		kessoku.Bind[oauth.Handler](kessoku.Provide(oauth.New)),
		kessoku.Bind[pixelfeddirectory.Handler](kessoku.Provide(pixelfeddirectory.New)),
		kessoku.Bind[statusedit.Handler](kessoku.Provide(statusedit.New)),
		kessoku.Bind[storyapiv1.Handler](kessoku.Provide(storyapiv1.New)),
		kessoku.Bind[story.Handler](kessoku.Provide(story.New)),
		kessoku.Bind[userappsettings.Handler](kessoku.Provide(userappsettings.New)),
		kessoku.Provide(handler.NewAPIHandlers),
	),
	kessoku.Provide(func(srv *http.Server, e *client.Event) *App { return &App{HTTPServer: srv, Router: e.Router} }),
)
