package admins

import (
	"errors"
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
	"strings"
)

type AdminsRoles struct {
	controllers.Base
}

func (c *AdminsRoles) Get() {
	var (
		adminId = c.Ctx.Params().GetIntDefault("adminId", 0)
		roles   []*model.Roles
		err     error
	)

	roles, err = service.NewAdminsRolesService().ListBy(adminId)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(roles)

}

func (c *AdminsRoles) PutBy(roles string) {
	var (
		adminId = c.Ctx.Params().GetInt64Default("adminId", 0)
		roleIds = strings.Split(roles, ",")
		num     int
		err     error
	)

	//管理员不允许修改
	if adminId < 1 {
		c.SendBadRequest(errors.New("参数错误"), nil)
		return
	}

	num, err = service.NewAdminsRolesService().Modify(adminId, roleIds, model.ORGANIZEID_OF_BACKEND)

	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(num)

}
