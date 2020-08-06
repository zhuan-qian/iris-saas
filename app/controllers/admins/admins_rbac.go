package admins

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
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
