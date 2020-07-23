package model

type MenusResources struct {
	MenuPath    string `json:"menuPath" xorm:"not null comment('菜单路径') unique('menus_resources_unique') CHAR(10)"`
	Obj         string `json:"obj" xorm:"not null comment('资源路径') unique('menus_resources_unique') CHAR(255)"`
	Act         string `json:"act" xorm:"not null comment('行为') unique('menus_resources_unique') CHAR(10)"`
	Genre       int8   `json:"genre" xorm:"not null comment('类型 0:管理端 1:组织端') unique('menus_resources_unique') TINYINT"`
	Description string `json:"description" xorm:"not null comment('资源说明') VARCHAR(127)"`

	Related *int8 `json:"related,omitempty" xorm:"-"`
}

func (m *MenusResources) TableName() string {
	return "menus_resources"
}
