package database

import (
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
	"strconv"
)

//数据库结构同步方法
func SyncDB() {
	db := common.GetDB()
	err := db.Sync2(
		new(model.SystemParameters),
		new(model.Users),
		new(model.Roles),
		new(model.Attachments),
		new(model.Locations),
		new(model.Admins),
		new(model.EsLostIndex),
		new(model.EsLostUpdate),
		new(model.Operations),
		new(model.Organizations),
		new(model.OrgGroups),
		new(model.Menus),
		new(model.RolesMenus),
		new(model.MenusResources),
	)
	if err != nil {
		panic("数据库结构同步失败 原因:" + err.Error())
	}

}

//基础数据构建
func BuildDBData() error {
	var (
		systemDao = dao.NewSystemParametersDao().WithSession(nil)
		err       error
	)

	//管理员创建
	if !systemDao.AddminAccountCreated() {
		_, err := service.NewAdminsService().Create("13012341234", "122333", "admin")
		if err != nil {
			return err
		}
		systemDao.Update(model.SYSTEM_PARAMETER_FOR_ADMIN_ACCOUNT, "1")
	}

	//管理员角色创建与关联
	systemDao.ConfigureEnableIfUnusedByInt(model.SYSTEM_PARAMETER_FOR_ADMIN_ROLE, "0", "role_for_admin",
		func() {
			m := &model.Roles{
				Name:   model.ROLES_NAME_OF_KING,
				OrgId:  model.ORGANIZEID_OF_BACKEND,
				Status: 1,
			}
			_, err = service.NewRolesService().Create(m)
			if err != nil {
				panic(err)
			}
			//casbin只有增加策略后角色与人员关联关系才可以被映射到,才允许用策略参数做判断比较
			_, err = common.GetCasbin().AddPolicy(model.ROLES_NAME_OF_KING, strconv.Itoa(model.ORGANIZEID_OF_BACKEND), "all-objs", "all-acts")
			if err != nil {
				return
			}
			_, err = service.NewAdminsRolesService().Modify(1, []string{strconv.Itoa(m.Id)}, model.ORGANIZEID_OF_BACKEND)
			if err != nil {
				panic(err)
			}
		})

	//城市地址库创建
	if !systemDao.LocationsBuilt() {
		sqlStr, err := common.ReadFileToString("toolkit/sql/locations.sql")
		if err != nil {
			return err
		}
		_, err = common.GetDB().Exec(sqlStr)
		if err != nil {
			return err
		}
		systemDao.Update(model.SYSTEM_PARAMETER_FOR_LOCATIONS_BUILT, "1")
	}

	//运营参数建立
	operationsBuild()

	//菜单构建
	systemDao.ConfigureEnableIfUnusedByInt(model.SYSTEM_PARAMETER_FOR_MENUS_BUILT, "0", "管理端与学校端菜单建立",
		func() {
			err = MenusBuild()
			if err != nil {
				panic(err)
			}
		})

	return nil
}

func operationsBuild() {
	o := dao.NewOperationsDao().WithSession(nil)
	o.InitIfNotExist(model.SYSTEM_PARAMETER_FOR_OPERATIONS_CRUMBS, "{}", "发现页面包屑导航")
	o.InitIfNotExist(model.SYSTEM_PARAMETER_FOR_OPERATIONS_GUIDE, "{}", "发现页指南路径")
	o.InitIfNotExist(model.SYSTEM_PARAMETER_FOR_OPERATIONS_SLIDE, "{}", "发现页轮播图")
	o.InitIfNotExist(model.SYSTEM_PARAMETER_FOR_OPERATIONS_HOTKEYWORD, "[]", "搜索页热门关键词")
}
