package dao

import (
	"fmt"
	"gold_hill/mine/model"
	"xorm.io/xorm"
)

type MenusResources struct {
	Base
}

func NewMenusResourcesDao() *MenusResources {
	return &MenusResources{}
}

func (d *MenusResources) WithSession(s *xorm.Session) *MenusResources {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *MenusResources) DeleteByMenuPathAndGenre(path string, genre int8) error {
	sql := fmt.Sprintf("delete from `%s` where `menuPath`='%s' and `genre`='%d'",
		new(model.MenusResources).TableName(), path, genre)
	_, err := d.session.Exec(sql)
	return err
}

func (d *MenusResources) ListBy(genre int8, menuPath *string) (err error, list []*model.MenusResources) {
	if menuPath != nil {
		d.session = d.session.Where("menuPath=?", *menuPath)
	}
	err = d.session.Where("genre=?", genre).Find(&list)
	return
}

func (d *MenusResources) ListByMenus(genre int8, menusPath []string) (err error, list []*model.MenusResources) {
	if menusPath != nil {
		d.session = d.session.In("menuPath", menusPath)
	}
	err = d.session.Where("genre=?", genre).Find(&list)
	return
}
