package admins

import (
	"encoding/json"
	"errors"
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/dao"
	"scaffold/model"
	"scaffold/service"
)

type RolesRbac struct {
	controllers.Base
}

func (c *RolesRbac) Get() {
	var (
		roleId, err = c.Ctx.Params().GetInt("roleId")
		list        []*model.RbacPolicy
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	list = service.NewRolesRbacService().List(model.ORGANIZEID_OF_BACKEND, roleId)

	c.SendSmile(list)
}

func (c *RolesRbac) Put() {
	var (
		roleId, err   = c.Ctx.Params().GetInt("roleId")
		p             = &params.RolesRbacPut{}
		policyListStr = c.Ctx.URLParam("policies")
		rolesDao      = dao.NewRolesDao().WithSession(nil)
		result        bool
	)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	p.RoleName, err = rolesDao.NameById(roleId)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	if policyListStr == "" {
		err = errors.New("policies参数缺失")
	}

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	err = json.Unmarshal([]byte(policyListStr), &p.PolicyList)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	err = c.Validate.Struct(*p)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	result, err = service.NewRolesRbacService().Modify(model.ORGANIZEID_OF_BACKEND, p.RoleName, p.PolicyList)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(result)
}
