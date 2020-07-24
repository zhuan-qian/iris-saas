package admins

import (
	"errors"
	"gold_hill/scaffold/app/controllers"
	"gold_hill/scaffold/app/controllers/params"
	"gold_hill/scaffold/common"
	"gold_hill/scaffold/dao"
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/service"
)

type RolesMenus struct {
	controllers.Base
}

func (c *RolesMenus) Put() {
	var (
		roleId, err = c.Ctx.Params().GetInt("roleId")
		p           = &params.RolesMenusPost{}
		result      bool
		roleName    string
	)

	if err != nil {
		err = errors.New("参数错误")
		goto REQUIRED_ERR
	}

	err = c.Ctx.ReadQuery(p)
	if err != nil {
		goto REQUIRED_ERR
	}

	err = c.Validate.Struct(*p)
	if err != nil {
		goto REQUIRED_ERR
	}

	roleName, err = dao.NewRolesDao().WithSession(nil).NameById(roleId)
	if err != nil {
		err = errors.New("参数错误")
		goto REQUIRED_ERR
	}

	result, err = service.NewRolesMenusService().Create(model.ORGANIZEID_OF_BACKEND, roleName, p.MenusPath)
	if err != nil {
		if common.IsRequireError(err) {
			goto REQUIRED_ERR
		}
		goto SERVER_ERR
	}
	c.SendSmile(result)
	return

REQUIRED_ERR:
	c.SendBadRequest(err.Error(), nil)
	return

SERVER_ERR:
	c.SendServerError(err.Error())
	return

}
