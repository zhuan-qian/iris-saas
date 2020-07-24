package admins

import (
	"gold_hill/scaffold/app/controllers"
	"gold_hill/scaffold/app/controllers/params"
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/service"
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
