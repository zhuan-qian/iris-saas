package client

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/app/controllers/params"
	"gold_hill/mine/common"
	"gold_hill/mine/service"
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
