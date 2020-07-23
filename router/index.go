package router

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

func Run(app *iris.Application) {

	//健康检查
	app.Get("/health", func(ctx iris.Context) {
		ctx.StatusCode(http.StatusOK)
		return
	})

	//注册应用程序接口
	AppRouter(app)

	//注册管理后台接口
	AdminRouter(app)

}
