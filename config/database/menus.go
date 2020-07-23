package database

import (
	"scaffold/dao"
	"scaffold/model"
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
		omenus = OrganizationMenus()
		err    error
	)

	err = MenusHandle(bmenus, model.MENU_GENRE_IS_BACKEND)
	if err != nil {
		return err
	}
	err = MenusHandle(omenus, model.MENU_GENRE_IS_ORGANIZATION)
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
			Title:       "店铺管理",
			Path:        "02000000",
			Route:       "/organization",
			Icon:        "peoples",
			Description: "对所有店铺的管理",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "店铺",
					Path:        "02010000",
					Route:       "list",
					Icon:        "",
					Description: "针对店铺查看与管理",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins/organizations",
							Act:         "get",
							Description: "店铺列表",
						},
						{
							Obj:         "/admins/organizations",
							Act:         "post",
							Description: "创建店铺",
						},
						{
							Obj:         "/admins/organizations/$v",
							Act:         "put",
							Description: "编辑店铺",
						},
					},
				},
				{
					Title:       "分组",
					Path:        "02020000",
					Route:       "group",
					Icon:        "",
					Description: "对店铺的分组",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        1,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins/orgGroups",
							Act:         "get",
							Description: "店铺分组列表",
						},
						{
							Obj:         "/admins/orgGroups",
							Act:         "post",
							Description: "创建店铺分组",
						},
						{
							Obj:         "/admins/orgGroups/$v",
							Act:         "put",
							Description: "编辑店铺分组",
						},
					},
				},
				//{
				//	Title:       "教师",
				//	Path:        "02030000",
				//	Route:       "teacher",
				//	Icon:        "",
				//	Description: "对学校教师的管理",
				//	Genre:       model.MENU_GENRE_IS_BACKEND,
				//	Sort:        1,
				//	Status:      1,
				//	Hidden:      1,
				//	Subs:        nil,
				//	Policies: []*model.RbacPolicyWithDescription{
				//		{
				//			Obj:         "/admins/teachers",
				//			Act:         "get",
				//			Description: "学校教师列表",
				//		},
				//		{
				//			Obj:         "/admins/teachers/$v",
				//			Act:         "put",
				//			Description: "编辑学校教师",
				//		},
				//	},
				//},
			},
		},
		{
			Title:       "会员管理",
			Path:        "03000000",
			Route:       "/user",
			Icon:        "user",
			Description: "对所有应用的会员进行统一管理",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "会员",
					Path:        "03010000",
					Route:       "list",
					Icon:        "",
					Description: "会员查看以及操作",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
					Policies: []*model.RbacPolicyWithDescription{
						{
							Obj:         "/admins/users",
							Act:         "get",
							Description: "会员列表",
						},
						{
							Obj:         "/admins/users",
							Act:         "post",
							Description: "创建会员",
						},
						{
							Obj:         "/admins/users/$v",
							Act:         "put",
							Description: "编辑会员",
						},
					},
				},
			},
		},
		{
			Title:       "商品管理",
			Path:        "04000000",
			Route:       "/goods",
			Icon:        "user",
			Description: "对所有应用的会员进行统一管理",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "商品",
					Path:        "04010000",
					Route:       "list",
					Icon:        "",
					Description: "管理组织商品列表",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
				{
					Title:       "子商品",
					Path:        "04020000",
					Route:       "subList",
					Icon:        "",
					Description: "管理组织子商品",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      1,
					Subs:        nil,
				},
				{
					Title:       "分类",
					Path:        "04020000",
					Route:       "category",
					Icon:        "",
					Description: "管理组织商品分类",
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
			Path:        "05000000",
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
					Path:        "05010000",
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
					Path:        "05020000",
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
					Path:        "05030000",
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
		{
			Title:       "财务管理",
			Path:        "06000000",
			Route:       "/finance",
			Icon:        "money",
			Description: "平台所有财务统一管理",
			Genre:       model.MENU_GENRE_IS_BACKEND,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "流水统计",
					Path:        "06010000",
					Route:       "list",
					Icon:        "",
					Description: "平台所有进出账流水统计",
					Genre:       model.MENU_GENRE_IS_BACKEND,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
	}
}

func OrganizationMenus() []*model.Menus {
	return []*model.Menus{
		{
			Title:       "首页管理",
			Path:        "01000000",
			Route:       "/",
			Icon:        "dashboard",
			Description: "平台所有数据统计管理",
			Genre:       model.MENU_GENRE_IS_ORGANIZATION,
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
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
		{
			Title:       "学生管理",
			Path:        "02000000",
			Route:       "/user",
			Icon:        "peoples",
			Description: "对本校所有学生进行管理",
			Genre:       model.MENU_GENRE_IS_ORGANIZATION,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "学生",
					Path:        "02010000",
					Route:       "list",
					Icon:        "",
					Description: "对本校所有学生的信息查看以及创建",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
		{
			Title:       "商品管理",
			Path:        "03000000",
			Route:       "/goods",
			Icon:        "shopping",
			Description: "对本校销售商品统一管理",
			Genre:       model.MENU_GENRE_IS_ORGANIZATION,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "商品",
					Path:        "03010000",
					Route:       "list",
					Icon:        "",
					Description: "管理本校所有商品",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        1,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
				{
					Title:       "子商品",
					Path:        "03020000",
					Route:       "subList",
					Icon:        "",
					Description: "管理本校所有子商品",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        1,
					Status:      1,
					Hidden:      1,
					Subs:        nil,
				},
				{
					Title:       "创建商品",
					Path:        "03030000",
					Route:       "create",
					Icon:        "",
					Description: "对商品的创建",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        1,
					Status:      1,
					Hidden:      1,
					Subs:        nil,
				},
				{
					Title:       "编辑商品",
					Path:        "03040000",
					Route:       "edit",
					Icon:        "",
					Description: "对商品的编辑",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        1,
					Status:      1,
					Hidden:      1,
					Subs:        nil,
				},
			},
		},
		{
			Title:       "运营管理",
			Path:        "04000000",
			Route:       "/operation",
			Icon:        "chart",
			Description: "管理运营相关操作",
			Genre:       model.MENU_GENRE_IS_ORGANIZATION,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "员工",
					Path:        "04010000",
					Route:       "user",
					Icon:        "",
					Description: "允许登录后台的员工",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
				{
					Title:       "角色",
					Path:        "04020000",
					Route:       "role",
					Icon:        "",
					Description: "对员工职能赋能管理",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        1,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
				{
					Title:       "权限",
					Path:        "04030000",
					Route:       "resource",
					Icon:        "",
					Description: "后台可用资源列表与描述",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        2,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
		{
			Title:       "财务管理",
			Path:        "05000000",
			Route:       "/finance",
			Icon:        "money",
			Description: "本校进出账财务查看与管理",
			Genre:       model.MENU_GENRE_IS_ORGANIZATION,
			Sort:        0,
			Status:      1,
			Hidden:      0,
			Subs: []*model.Menus{
				{
					Title:       "流水统计",
					Path:        "05010000",
					Route:       "list",
					Icon:        "",
					Description: "本校所有进出账流水统计",
					Genre:       model.MENU_GENRE_IS_ORGANIZATION,
					Sort:        0,
					Status:      1,
					Hidden:      0,
					Subs:        nil,
				},
			},
		},
	}
}
