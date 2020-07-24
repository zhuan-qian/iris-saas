package config

import (
	"gold_hill/scaffold/common"
	"gold_hill/scaffold/config/database"
	"gold_hill/scaffold/config/directory"
	"gold_hill/scaffold/service/payment"
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
