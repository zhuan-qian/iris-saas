package resource

import "github.com/kataras/iris/v12"

func Run(app *iris.Application) {
	app.RegisterView(iris.HTML("./public", ".html"))
	app.HandleDir("/static", "public/static")
	//app.HandleDir("/admin", "public/admin")
	//app.HandleDir("/courses", "public/resource/courses")
}
