package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/app/controllers/params"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
)

type Users struct {
	controllers.Base
}

//获取用户列表
func (c *Users) Get() {
	var (
		p     = new(params.UsersGet)
		m     []*model.Users
		count int64
		err   = c.Ctx.ReadQuery(p)
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	m, count, err = service.NewUsersService().GetListByKey(p.Limit, p.Page, p.Keyword)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(controllers.CompactListAndCount(m, count))
}

func (c *Users) GetBy(id int64) {
	var (
		m   = &model.Users{}
		err error
	)
	m, err = service.NewUsersService().InfoBy(id)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(m)
}

//注册用户
func (c *Users) Post() {
	var (
		m   = &model.Users{}
		err = c.Ctx.ReadForm(m)
		num int64
	)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	err = c.Validate.Struct(*m)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	num, err = service.NewUsersService().Create(m)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(num)
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
