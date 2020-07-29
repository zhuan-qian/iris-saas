package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/app/controllers/params"
	"gold_hill/mine/dao"
	"gold_hill/mine/service"
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

func (c *Operations) Post() {
	var (
		p   = &params.OperationsPost{}
		err = c.Ctx.ReadForm(p)
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	err = c.Validate.Struct(*p)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	err = service.NewOperationsService().Create(p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)
	return
}

func (c *Operations) PutBy(name string) {
	var (
		p   = &params.OperationsPut{}
		err = c.Ctx.ReadQuery(p)
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	err = c.Validate.Struct(*p)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	err = service.NewOperationsService().ModifyBy(name, p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)
	return

}

func (c *Operations) DeleteBy(name string) {

	err := service.NewOperationsService().DeleteBy(name)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)

}
