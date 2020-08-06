package admins

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
)

type OrgGroups struct {
	controllers.Base
}

//获取学校分组列表
func (c *OrgGroups) Get() {
	var (
		p     = new(params.OrgGroupsGet)
		m     []*model.OrgGroups
		count int64
		err   = c.Ctx.ReadQuery(p)
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

	m, count, err = service.NewOrgGroupsService().GetListByKey(p.Limit, p.Page)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(controllers.CompactListAndCount(m, count))
}

func (c *OrgGroups) Post() {
	var (
		m   = &model.OrgGroups{}
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

	num, err = service.NewOrgGroupsService().Create(m)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(num)
}

func (c *OrgGroups) PutBy(id int) {
	var (
		p   = &params.OrgGroupsPut{}
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

	_, err = service.NewOrgGroupsService().Modify(id, p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	err = service.NewOrganizationsService().ModifyBy(id)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)

}
