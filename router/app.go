package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gold_hill/scaffold/app/controllers/client"
	"gold_hill/scaffold/app/middleware"
)

func AppRouter(app *iris.Application) {

	mvc.Configure(app.Party("/v1"), func(m *mvc.Application) {
		var (
			unauth   = m.Party("/").Register(validator.New())
			needauth = m.Party("/", middleware.UsersVerify).Register(validator.New())
		)

		//注册或登录
		unauth.Party("/users").Handle(new(client.UsersAccess))
		unauth.Party("/sms").Handle(new(client.Sms))
		unauth.Party("/operations").Handle(new(client.Operations))
		needauth.Party("/users").Handle(new(client.Users))
	})
}
