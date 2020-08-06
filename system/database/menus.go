package database

import (
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"strings"
)

func MenusBuild() error {
	var (
		err error
	)

	err = Background()
	if err != nil {
		return err
	}

	return nil
}

func Background() error {
	var (
		bmenus = BackgroundMenus()
		err    error
	)

	err = MenusHandle(bmenus, model.MENU_GENRE_IS_BACKEND)
	if err != nil {
		return err
	}
	return err
}

func MenusHandle(list []*model.Menus, domain int8) error {
	var (
		menuDao = dao.NewMenusDao().WithSession(nil)
		mrDao   = dao.NewMenusResourcesDao().WithSession(nil)
		err     error
	)
	for _, m := range list {
		//插入菜单并更新path
		err = menuDao.DeleteBy(m.Path, m.Genre)
		if err != nil {
			return err
		}
		_, err = menuDao.InsertOne(m)
		if err != nil {
			return err
		}

		//插入菜单与资源关联关系
		err = mrDao.DeleteByMenuPathAndGenre(m.Path, domain)
		if err != nil {
			return err
		}
		var mrList []*model.MenusResources
		for _, v := range m.Policies {
			mrList = append(mrList, &model.MenusResources{
				MenuPath:    m.Path,
				Obj:         strings.ToLower(v.Obj),
				Act:         strings.ToLower(v.Act),
				Genre:       domain,
				Description: v.Description,
			})
		}
		if mrList != nil {
			_, err := mrDao.Insert(&mrList)
			if err != nil {
				return err
			}
		}

		if m.Subs != nil {
			err = MenusHandle(m.Subs, domain)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func BackgroundMenus() []*model.Menus {
	return []*model.Menus{
		{
			Title:       "首页管理",
			Path:        "01000000",
			Route:       "/",
			Icon:        "dashboard",
			Description: "平台所有数据统计管理",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "首页",
					Path:        "01010000",
					Route:       "dashboard",
					Icon:        "",
					Description: "平台所有数据统计",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
		{
			Title:       "运营管理",
			Path:        "02000000",
			Route:       "/operation",
			Icon:        "chart",
			Description: "管理运营相关操作",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "员工",
					Path:        "02010000",
					Route:       "user",
					Icon:        "",
					Description: "允许登录后台的员工",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins",
							Act:         "get",
							Description: "员工列表",
						},
						{
							Obj:         "/admins",
							Act:         "post",
							Description: "创建员工",
						},
						{
							Obj:         "/admins/$v",
							Act:         "put",
							Description: "编辑员工",
						},
						{
							Obj:         "/admins/$v/roles/$v",
							Act:         "put",
							Description: "员工关联角色",
						},
					},
				},
				{
					Title:       "角色",
					Path:        "02020000",
					Route:       "role",
					Icon:        "",
					Description: "对员工职能赋能管理",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins/roles",
							Act:         "get",
							Description: "角色列表",
						},
						{
							Obj:         "/admins/roles",
							Act:         "post",
							Description: "创建角色",
						},
						{
							Obj:         "/admins/roles/$v",
							Act:         "put",
							Description: "编辑角色",
						},
						{
							Obj:         "/admins/roles/$v/rbac",
							Act:         "put",
							Description: "角色关联权限",
						},
					},
				},
				{
					Title:       "权限",
					Path:        "02030000",
					Route:       "resource",
					Icon:        "",
					Description: "后台可用权限列表与描述",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins/rbac",
							Act:         "get",
							Description: "权限列表",
						},
					},
				},
			},
		},
	}
}
