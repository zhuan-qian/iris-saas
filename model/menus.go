package model

const (
	MENU_GENRE_IS_BACKEND      = 0
	MENU_GENRE_IS_ORGANIZATION = 1
	SUFFIX_OF_LEVEL_0          = "00000000"
	SUFFIX_OF_LEVEL_1          = "000000"
	SUFFIX_OF_LEVEL_2          = "0000"
	SUFFIX_OF_LEVEL_3          = "00"
)

type Menus struct {
	Path        string `json:"path" xorm:"not null comment('父级id') CHAR(10) unique('menus_unique_path_genre')"`
	Route       string `json:"route" xorm:"not null comment('菜单路由') CHAR(255)"`
	Title       string `json:"title" xorm:"not null comment('菜单标题') CHAR(32)"`
	Icon        string `json:"icon" xorm:"not null comment('菜单图标字符串')"`
	Description string `json:"description" xorm:"null comment('说明') CHAR(255)"`
	Genre       int8   `json:"genre" xorm:"not null comment('类型 0:管理端 1:组织端') unique('menus_unique_path_genre') TINYINT"`
	Sort        int8   `json:"sort" xorm:"not null default 0 comment('同级排序') TINYINT"`
	Status      int8   `json:"status" xorm:"not null comment('状态 0:禁用 1:启用') TINYINT"`
	Hidden      int8   `json:"hidden" xorm:"not null comment('是否隐藏菜单 0: 显示 1: 隐藏') TINYINT"`

	Related   *int8                        `json:"related,omitempty" xorm:"-"`
	Subs      []*Menus                     `json:"subs,omitempty" xorm:"-"`
	Policies  []*RbacPolicyWithDescription `json:"policies,omitempty" xorm:"-"`
	Resources []*MenusResources            `json:"resources,omitempty" xorm:"-"`
}

func (m *Menus) TableName() string {
	return "menus"
}
