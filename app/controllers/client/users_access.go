package client

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/service"
)

type UsersAccess struct {
	controllers.Base
}

//创建账号或登录账号
func (c *UsersAccess) Post() {
	var (
		user  = &params.UsersPost{}
		err   = c.Ctx.ReadJSON(user)
		token string
	)
	if err != nil {
		c.SendBadRequest("参数缺失:"+err.Error(), nil)
		return
	}

	err = c.Validate.Struct(*user)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	token, err = service.NewUsersService().LoginOrRegister(user.Account, user.Code)
	if err != nil {
		if common.IsRequireError(err) {
			c.SendBadRequest(err.Error(), nil)
			return
		}
		c.SendServerError(err.Error())
		return
	}
	c.SendCreated(token)
	return
}
