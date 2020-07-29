package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
)

type AdminsRbac struct {
	controllers.Base
}

func (c *AdminsRbac) Get() {
	var (
		adminId, err = c.Ctx.Params().GetInt("adminId")
		list         []*model.RbacPolicy
	)

	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}

	list, err = service.NewAdminsRbacService().List(adminId)
	if err != nil {
		c.SendServerError(err)
		return
	}

	c.SendSmile(list)
}
