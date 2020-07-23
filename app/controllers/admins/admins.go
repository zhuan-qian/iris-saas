package admins

import (
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/model"
	"scaffold/service"
)

type Admins struct {
	controllers.Base
}

//获取管理员列表
func (c *Admins) Get() {
	var (
		p     = new(params.AdminsGet)
		m     []*model.Admins
		count int64
		err   error
	)

	err = c.Ctx.ReadQuery(p)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	m, count, err = service.NewAdminsService().GetListByKey(p.Limit, p.Page, p.Keyword)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(controllers.CompactListAndCount(m, count))
}

//注册员工
func (c *Admins) Post() {
	var (
		m   = &params.AdminsPost{}
		err = c.Ctx.ReadForm(m)
		id  int
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

	id, err = service.NewAdminsService().Create(m.Account, m.Password, m.Nickname)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(id)
}

func (c *Admins) PutBy(id int) {
	var (
		p   = &params.AdminsModify{}
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
	ok, err = service.NewAdminsService().Modify(id, p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(ok)
}
