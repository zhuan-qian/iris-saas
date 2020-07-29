package client

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/dao"
)

type Operations struct {
	controllers.Base
}

func (c *Operations) Get() {
	result := dao.NewOperationsDao().WithSession(nil).Get()

	c.SendSmile(result)
	return
}

func (c *Operations) GetBy(name string) {
	result := dao.NewOperationsDao().WithSession(nil).GetByName(name)

	c.SendSmile(result)
	return
}
