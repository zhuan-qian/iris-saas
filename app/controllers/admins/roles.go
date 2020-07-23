package admins

import (
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/common"
	"scaffold/model"
	"scaffold/service"
)

type Roles struct {
	controllers.Base
}

func (c *Roles) Get() {
	var (
		p     = new(params.RolesGet)
		m     []model.Roles
		count int64
		err   = c.Ctx.ReadQuery(p)
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	m, count, err = service.NewRolesService().GetListByKey(model.ORGANIZEID_OF_BACKEND, p.Limit, p.Page, p.Keyword)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(controllers.CompactListAndCount(m, count))
}

func (c *Roles) Post() {
	var (
		m   = &model.Roles{}
		err = c.Ctx.ReadForm(m)
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

	_, err = service.NewRolesService().Create(m)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(m.Id)
}

func (c *Roles) PutBy(id int) {
	var (
		p   = &params.RolesPut{}
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

	_, err = service.NewRolesService().Modify(model.ORGANIZEID_OF_BACKEND, id, p)
	if err != nil {
		if common.IsRequireError(err) {
			c.SendBadRequest(err.Error(), nil)
			return
		}
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)

}
