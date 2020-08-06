package admins

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
)

type AdminsMenus struct {
	controllers.Base
}

func (c *AdminsMenus) Get() {
	var (
		p    = &params.MenusGet{}
		list []*model.Menus
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

	list, err = service.NewMenusService().List(model.ORGANIZEID_OF_BACKEND, p, model.MENU_GENRE_IS_BACKEND)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(list)

}
