package service

import (
	"scaffold/app/controllers/params"
	"scaffold/common"
	"scaffold/dao"
	"scaffold/model"
	"strings"
)

type Menus struct {
	d *dao.Menus
}

func NewMenusService() *Menus {
	return &Menus{d: dao.NewMenusDao().WithSession(nil)}
}

func (s *Menus) List(orgId int, p *params.MenusGet, genre int8) (m []*model.Menus, err error) {
	var (
		menus          []string
		menusMap       = make(map[string]bool)
		rolesName      []string
		rmDao          = dao.NewRolesMenusDao().WithSession(nil)
		rolesMenusList []*model.RolesMenus
		isKing         bool
	)

	//如果有传角色参数 则取出角色相关菜单列表
	if p.RolesName != nil {
		rolesName = strings.Split(strings.Trim(strings.TrimSpace(*p.RolesName), ","), ",")
		isKing = dao.NewRolesDao().ContainsKing(rolesName)
		if isKing {
			goto ADMIN_AIRBORNE
		}
		rolesMenusList, err = rmDao.ListByRolesName(orgId, rolesName)
		if err != nil {
			return nil, err
		}

		//如果角色没有关联任何菜单 则直接返回
		if p.OnlyRelated != nil && *p.OnlyRelated == 1 {
			if rolesMenusList == nil {
				return
			}
			menus = rmDao.CollectMenuBy(rolesMenusList)
		}
	}

ADMIN_AIRBORNE:

	//获取菜单列表
	m, err = s.d.ListAndDescBy(genre, menus)

	if err != nil || len(m) == 0 {
		return nil, err
	}

	//标记已有关联菜单
	if p.TagRelated != nil && *p.TagRelated == 1 {
		menusMap = rmDao.MapMenuBy(rolesMenusList)

		for i, v := range m {
			_, ok := menusMap[v.Path]

			if ok || isKing {
				m[i].Related = common.Int8Ptr(1)
			}
		}
	}

	//获取菜单相关资源列表
	err = s.fillResources(orgId, genre, rolesName, m)
	if err != nil {
		return nil, err
	}

	if p.IsTree != 1 {
		return
	}

	//树状结构建立
	m = s.d.BuildTreeBy(m)

	//按path排序 从小喝到大
	s.d.Sort(true, m)

	return
}

func (s *Menus) fillResources(orgId int, genre int8, rolesName []string, list []*model.Menus) error {
	var (
		mrList []*model.MenusResources

		rbacDao          = dao.NewRbacDao()
		rbacs            model.RbacPolicies
		objsOfRolesMap   = make(map[string]bool)
		menusResourceMap = make(map[string][]*model.MenusResources)
		menus            []string
		err              error
	)

	menus = dao.NewMenusDao().CollectPathBy(list)

	//获取菜单相关资源列表
	err, mrList = dao.NewMenusResourcesDao().WithSession(nil).ListByMenus(genre, menus)
	if err != nil {
		return err
	}

	//获取角色的资源map
	if rolesName != nil {
		rbacs = rbacDao.CollectRolesPolicies(orgId, rolesName)
		objsOfRolesMap = rbacDao.ParsePoliciesToObjsAndActsMap(rbacs)
	}

	//为菜单资源打标签并按菜单分组
	for i, v := range mrList {
		//角色具备菜单的资源 并且允许打标签
		if _, ok := objsOfRolesMap[v.Obj+","+v.Act]; ok {
			mrList[i].Related = common.Int8Ptr(1)
		}

		//显示所有资源
		if _, ok := menusResourceMap[v.MenuPath]; ok {
			menusResourceMap[v.MenuPath] = append(menusResourceMap[v.MenuPath], mrList[i])
		} else {
			menusResourceMap[v.MenuPath] = []*model.MenusResources{mrList[i]}
		}
	}

	//将资源置入菜单
	for i, v := range list {
		if _, ok := menusResourceMap[v.Path]; ok {
			list[i].Resources = menusResourceMap[v.Path]
		}
	}
	return nil
}
