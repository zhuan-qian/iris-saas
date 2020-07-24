package client

import (
	"gold_hill/scaffold/app/controllers"
	"gold_hill/scaffold/app/controllers/params"
	"gold_hill/scaffold/dao"
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/service"
)

type Users struct {
	controllers.Base
}

func (c *Users) Get() {
	var (
		userCacheInfo = c.Ctx.Values().Get("user").(*model.Users)
		user          = &model.Users{}
		err           error
	)

	user, err = dao.NewUsersDao().WithSession(nil).Info(userCacheInfo.Id)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(user)
}

func (c *Users) PutBy(id int64) {
	var (
		p   = &params.UsersModify{}
		err = c.Ctx.ReadQuery(p)
		ok  bool
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
	ok, err = service.NewUsersService().Modify(id, p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(ok)
}
