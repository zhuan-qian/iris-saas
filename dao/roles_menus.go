package dao

import (
	"fmt"
	"gold_hill/scaffold/model"
	"xorm.io/xorm"
)

type RolesMenus struct {
	Base
}

func NewRolesMenusDao() *RolesMenus {
	return &RolesMenus{}
}

func (d *RolesMenus) WithSession(s *xorm.Session) *RolesMenus {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *RolesMenus) listByRoleName(orgId int, roleName string) (list []*model.RolesMenus, err error) {
	err = d.session.Where("orgId=?", orgId).And("roleName=?", roleName).Find(&list)
	return
}

func (d *RolesMenus) listByRolesName(orgId int, rolesName []string) (list []*model.RolesMenus, err error) {
	err = d.session.Where("orgId=?", orgId).In("roleName", rolesName).Find(&list)
	return
}

func (d *RolesMenus) ListByRolesName(orgId int, rolesName []string) (list []*model.RolesMenus, err error) {
	var (
		rl = len(rolesName)
	)
	switch {
	case rl > 1:
		return d.listByRolesName(orgId, rolesName)

	case rl == 1:
		return d.listByRoleName(orgId, rolesName[0])

	default:
		return
	}
}

func (d *RolesMenus) CollectMenuBy(list []*model.RolesMenus) (menus []string) {
	for _, v := range list {
		menus = append(menus, v.MenuPath)
	}
	return
}

func (d *RolesMenus) MapMenuBy(list []*model.RolesMenus) (menus map[string]bool) {
	menus = make(map[string]bool)
	for _, v := range list {
		menus[v.MenuPath] = true
	}
	return
}

func (d *RolesMenus) DeleteByRoleName(orgId int, roleName string) error {
	sql := fmt.Sprintf("delete from `%s` where orgId = %d and roleName='%s'", new(model.RolesMenus).TableName(), orgId, roleName)
	_, err := d.session.Exec(sql)
	return err
}
