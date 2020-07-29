package dao

import (
	"gold_hill/mine/model"
	"xorm.io/xorm"
)

type Operations struct {
	Base
}

func NewOperationsDao() *Operations {
	return &Operations{}
}

func (d *Operations) WithSession(s *xorm.Session) *Operations {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Operations) Get() []model.Operations {
	var (
		m    []model.Operations
		list = []string{
			model.SYSTEM_PARAMETER_FOR_OPERATIONS_CRUMBS,
			model.SYSTEM_PARAMETER_FOR_OPERATIONS_GUIDE,
			model.SYSTEM_PARAMETER_FOR_OPERATIONS_SLIDE,
			model.SYSTEM_PARAMETER_FOR_OPERATIONS_HOTKEYWORD,
		}
	)
	d.session.In("name", list).Find(&m)
	return m
}

func (d *Operations) GetByName(name string) string {
	m := &model.Operations{}
	d.session.Cols("params").Where("name=?", name).Get(m)

	return m.Params
}

//配置是否已存在查询
func (d *Operations) InitIfNotExist(name string, defaultValue string, description string) bool {
	m := &model.Operations{}
	find, _ := d.session.Where("name=?", name).Exist(m)
	if !find {
		m.Name = name
		m.Params = defaultValue
		m.Description = description
		d.InsertOne(m)
		return true
	}
	return true
}
