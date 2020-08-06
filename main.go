package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
	"zhuan-qian/go-saas/system"
	"os"
)

func main() {
	var (
		app *iris.Application
	)

	//系统构建
	system.BaseBuild()
	app=system.AppBuild()

	//启动应用 尽可能的优化性能
	//app.Run(iris.TLS("sports.eastday.com:443","/var/www/server.crt","/var/www/server_private.key"),iris.WithOptimizations)
	app.Run(iris.Addr(":"+os.Getenv("APP_PORT")), iris.WithOptimizations)

}
