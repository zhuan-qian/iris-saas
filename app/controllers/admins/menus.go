package admins

import (
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/model"
	"scaffold/service"
)

type Menus struct {
	controllers.Base
}

func (c *Menus) Get() {
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
