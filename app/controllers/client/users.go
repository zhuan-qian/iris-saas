package client

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
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
