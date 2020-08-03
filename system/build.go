package system

import (
	"context"
	"github.com/iris-contrib/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"gold_hill/mine/app/commands/crons"
	"gold_hill/mine/common"
	"gold_hill/mine/router"
	"gold_hill/mine/service/log"
	"gold_hill/mine/system/database"
	"gold_hill/mine/system/directory"
	"gold_hill/mine/system/resource"
	"time"
)

//系统依赖构建
func BaseBuild() {
	var (
		err error
	)

	//环境配置获取
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	//必备第三方组件错误校验
	//日志加载
	log.Get()

	//rbac基础数据表构建
	common.GetCasbin()
	//payment.GetAliPayHandle()
	//payment.GetWxPayHandle()

	//数据初始化
	err = database.Run()
	if err != nil {
		panic(err)
	}

	//目录初始化
	err = directory.InitDirectory()
	if err != nil {
		panic(err)
	}

	//启动定时任务
	crons.Run()
}

func AppBuild() *iris.Application {
	app := iris.New()

	//当panic时将状态恢复并将状态写为500
	app.Use(recover.New())

	//测试环境开启跨域 方便测试
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
	app.Use(crs)

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

	//静态资源
	resource.Run(app)

	//构建路由
	router.Run(app)

	return app
}
