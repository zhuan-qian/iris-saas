package admins

import (
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/model"
	"scaffold/service"
)

type Organizations struct {
	controllers.Base
}

//获取学校列表
func (c *Organizations) Get() {
	var (
		p     = new(params.OrganizationsGet)
		m     []*model.Organizations
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

	m, count, err = service.NewOrganizationsService().GetListByKey(p.Limit, p.Page, p.Keyword, p.GroupId, p.Sort, p.SortType)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(controllers.CompactListAndCount(m, count))
}

func (c *Organizations) GetBy(id int64) {
	var (
		m   = &model.Organizations{}
		err error
	)
	m, err = service.NewOrganizationsService().InfoBy(id)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(m)
}

func (c *Organizations) Post() {
	//var (
	//	p       = &params.OrganizationsPost{}
	//	err     = c.Ctx.ReadForm(p)
	//	orgId   int
	//	adminId int
	//)
	//if err != nil {
	//	c.SendBadRequest(err.Error(), nil)
	//	return
	//}
	//err = c.Validate.Struct(*p)
	//if err != nil {
	//	c.SendBadRequest(err.Error(), nil)
	//	return
	//}
	//
	//orgId, err = service.NewOrganizationsService().Create(p)
	//if err != nil {
	//	c.SendServerError(err.Error())
	//	return
	//}
	//adminId, err = service.NewWorkersService().Create(orgId, *p.Phone, *p.Password, "admin")
	//if err != nil {
	//	c.SendServerError(err.Error())
	//	return
	//}
	//m := &model.Roles{
	//	Name:   model.ROLES_NAME_OF_KING,
	//	OrgId:  orgId,
	//	Status: 1,
	//}
	//_, err = service.NewRolesService().Create(m)
	//if err != nil {
	//	c.SendServerError(err.Error())
	//	return
	//}
	//_, err = common.GetCasbin().AddPolicy(model.ROLES_NAME_OF_KING, strconv.Itoa(orgId), "all-objs", "all-acts")
	//if err != nil {
	//	c.SendServerError(err.Error())
	//	return
	//}
	//_, err = service.NewAdminsRolesService().Modify(adminId, []string{strconv.Itoa(m.Id)}, orgId)
	//if err != nil {
	//	c.SendServerError(err.Error())
	//	return
	//}
	//
	//c.SendCreated(true)
}

func (c *Organizations) PutBy(id int) {
	var (
		p   = &params.OrganizationsPut{}
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

	_, err = service.NewOrganizationsService().Modify(id, p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)

}
