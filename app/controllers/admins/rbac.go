package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/app/controllers/params"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
)

type Rbac struct {
	controllers.Base
}

//获取资源列表
func (c *Rbac) Get() {
	var (
		p    = &params.RbacGet{}
		list []*model.MenusResources
		err  = c.Ctx.ReadQuery(p)
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

	err, list = service.NewRbacService().PoliciesBy(model.MENU_GENRE_IS_BACKEND, p.MenuPath)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(list)
}
