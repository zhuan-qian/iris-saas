package admins

import (
	"scaffold/app/controllers"
	"scaffold/app/controllers/params"
	"scaffold/common"
	"scaffold/dao"
	"scaffold/model"
	"scaffold/service"
	"strings"
)

type Info struct {
	controllers.Base
}

//获取员工菜单与权限
func (c *Info) Get() {
	var (
		admin     = c.Ctx.Values().Get(dao.KEY_FOR_ADMIN_INFO).(*model.Admins)
		rbacDao   = dao.NewRbacDao()
		rolesName = rbacDao.GetRolesNameBy(model.ORGANIZEID_OF_BACKEND, admin.Id)
		p         = &params.MenusGet{
			IsTree:      1,
			RolesName:   common.StringPtr(strings.Join(rolesName, ",")),
			OnlyRelated: common.Int8Ptr(1),
		}
		list, err = service.NewMenusService().List(model.ORGANIZEID_OF_BACKEND, p, model.MENU_GENRE_IS_BACKEND)
		result    = make(map[string]interface{})
	)
	if err != nil {
		c.SendBadRequest(err.Error(), nil)
		return
	}
	result["menus"] = list
	result["info"] = admin
	result["roles"] = rolesName
	c.SendSmile(result)
}
