package service

import (
	"gold_hill/mine/common"
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
	"strings"
)

type RolesMenus struct {
	d *dao.RolesMenus
}

func NewRolesMenusService() *RolesMenus {
	return &RolesMenus{d: dao.NewRolesMenusDao().WithSession(nil)}
}

func (s *RolesMenus) Create(orgId int, roleName string, menuPath string) (result bool, err error) {
	var (
		menuDao   = dao.NewMenusDao().WithSession(nil)
		menusPath []string
		menus     []*model.Menus
		relations []*model.RolesMenus
	)

	menuPath = common.TrimCommaAndSpace(menuPath)
	menusPath = strings.Split(menuPath, ",")

	menus, err = menuDao.ListBy(model.MENU_GENRE_IS_BACKEND, menusPath)
	if err != nil {
		return false, err
	}
	if menus == nil {
		return false, common.NewRequireError("菜单均无效")
	}

	menusPath = menuDao.CollectPathBy(menus)

	err = s.d.DeleteByRoleName(orgId, roleName)
	if err != nil {
		return false, err
	}

	for _, v := range menusPath {
		relations = append(relations, &model.RolesMenus{
			OrgId:    orgId,
			RoleName: roleName,
			MenuPath: v,
		})
	}

	_, err = s.d.Insert(&relations)
	if err != nil {
		return false, err
	}
	return true, nil
}
