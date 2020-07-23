package config

import (
	"scaffold/common"
	"scaffold/config/database"
	"scaffold/config/directory"
	"scaffold/service/payment"
)

func Build() {
	var (
		err error
	)

	//必备第三方组件错误校验

	//rbac基础数据表构建
	common.GetCasbin()
	payment.GetAliPayHandle()
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

}
