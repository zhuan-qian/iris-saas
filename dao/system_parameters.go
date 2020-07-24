package dao

import (
	"gold_hill/scaffold/model"
	"xorm.io/xorm"
)

type SystemParameters struct {
	Base
}

func NewSystemParametersDao() *SystemParameters {
	return &SystemParameters{}
}

func (d *SystemParameters) WithSession(s *xorm.Session) *SystemParameters {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

//管理员是否已创建
func (d *SystemParameters) AddminAccountCreated() bool {
	m := &model.SystemParameters{}
	find, _ := d.session.Cols("config").Where("name=?", model.SYSTEM_PARAMETER_FOR_ADMIN_ACCOUNT).Get(m)
	if !find {
		d.Add(model.SYSTEM_PARAMETER_FOR_ADMIN_ACCOUNT, "0", "管理员账号是否已创建")
		return false
	}
	if m.Config == "0" {
		return false
	}
	return true
}

//全球城市信息是否已构建
func (d *SystemParameters) LocationsBuilt() bool {
	m := &model.SystemParameters{}
	find, _ := d.session.Cols("config").Where("name=?", model.SYSTEM_PARAMETER_FOR_LOCATIONS_BUILT).Get(m)
	if !find {
		d.Add(model.SYSTEM_PARAMETER_FOR_LOCATIONS_BUILT, "0", "全球城市信息是否已构建")
		return false
	}
	if m.Config == "0" {
		return false
	}
	return true
}

//通用配置参数构建
func (d *SystemParameters) ConfigureEnableIfUnusedByInt(configName string, defaultValue string, description string,
	buildFunc func()) bool {

	m := &model.SystemParameters{}
	find, _ := d.session.Cols("config").Where("name=?", configName).Get(m)
	if !find {
		d.Add(configName, defaultValue, description)
	}
	if m.Config == "1" {
		return true
	}
	buildFunc()
	d.Update(configName, "1")
	return true
}

//系统参数增加
func (d *SystemParameters) Add(name string, config string, description string) bool {
	m := &model.SystemParameters{
		Name:        name,
		Config:      config,
		Description: description,
	}
	num, _ := d.InsertOne(m)
	if num > 0 {
		return true
	}
	return false
}

//系统参数更新
func (d *SystemParameters) Update(name string, config string) bool {
	m := &model.SystemParameters{
		Config: config,
	}
	num, err := d.session.Cols("config").Where("name=?", name).Update(m)
	if err != nil {
		panic(err)
	}
	if num > 0 {
		return true
	}
	return false

}
