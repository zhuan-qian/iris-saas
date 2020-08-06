package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"zhuan-qian/go-saas/app/controllers/admins"
	"zhuan-qian/go-saas/app/middleware"
)

func AdminRouter(app *iris.Application) {

	app.Options("/{any:path}", func(ctx iris.Context) {})
	mvc.Configure(app.Party("/admins/token"), func(m *mvc.Application) {
		m.Handle(new(admins.AdminsToken))
	})

	mvc.Configure(app.Party("/admins"), func(m *mvc.Application) {
		m.Router.Use(middleware.AdminsVerify)
		m.Register(validator.New())
		m.Party("/info").Handle(new(admins.Info))
	})
	mvc.Configure(app.Party("/admins"), func(m *mvc.Application) {
		m.Router.Use(middleware.AdminsVerify, middleware.AdminRbacVerify)
		m.Register(validator.New())
		m.Party("/").Handle(new(admins.Admins))
		m.Party("/{adminId:int min(1)}/roles").Handle(new(admins.AdminsRoles))
		m.Party("/{adminId:int min(1)}/rbac").Handle(new(admins.AdminsRbac))
		m.Party("/rbac").Handle(new(admins.Rbac))
		m.Party("/menus").Handle(new(admins.Menus))
		m.Party("/roles").Handle(new(admins.Roles))
		m.Party("/roles/{roleId:int min(1)}/rbac").Handle(new(admins.RolesRbac))
		m.Party("/roles/{roleId:int min(1)}/menus").Handle(new(admins.RolesMenus))
		m.Party("/users").Handle(new(admins.Users))
		m.Party("/attachments", iris.LimitRequestBodySize(100<<20)).Handle(new(admins.Attachments))
		m.Party("/operations").Handle(new(admins.Operations))
		m.Party("/organizations").Handle(new(admins.Organizations))
		m.Party("/orgGroups").Handle(new(admins.OrgGroups))
		m.Party("/locations").Handle(new(admins.Locations))
	})

}
