package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"os"
	"gold_hill/scaffold/app/commands/crons"
	"gold_hill/scaffold/common"
	"gold_hill/scaffold/config"
	"gold_hill/scaffold/config/resource"
	"gold_hill/scaffold/router"
	"time"
)

func main() {
	var (
		app *iris.Application
		err error
	)

	//环境配置获取
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	app = iris.New()

	//当panic时将状态恢复并将状态写为500
	app.Use(recover.New())

	app.Logger().SetLevel("info")
	//app.Logger().Handle()

	if common.IsDebug() {
		//测试环境开启跨域 方便测试
		crs := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
			AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
			AllowCredentials: true,
		})
		app.Use(crs)
	} else {

		//TODO 后期嵌入返回值中
		//日志服务
		f := common.LogFile()
		defer f.Close()
		app.Logger().SetOutput(f)
		requestLogger := logger.New(logger.Config{Status: true, IP: true, Method: true, Path: true, Query: true,
			MessageContextKeys: []string{"logger_message"},
			MessageHeaderKeys:  []string{"User-Agent"},
		})
		app.Use(requestLogger)

		//优雅关闭
		iris.RegisterOnInterrupt(func() {
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		})
		app.ConfigureHost(func(h *iris.Supervisor) {
			//注册关闭服务
			h.RegisterOnShutdown(func() {
				println("服务已关闭")
			})

			//注册错误服务
			h.RegisterOnError(func(err error) {
				//To do anything...
			})
		})

	}

	//系统构建
	config.Build()

	//静态资源
	resource.Run(app)

	//构建路由
	router.Run(app)

	//启动定时任务
	crons.Run()

	//启动应用 尽可能的优化性能
	//app.Run(iris.TLS("sports.eastday.com:443","/var/www/server.crt","/var/www/server_private.key"),iris.WithOptimizations)
	app.Run(iris.Addr(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))), iris.WithOptimizations)

}
