package model

type RolesMenus struct {
	OrgId    int    `json:"orgId" xorm:"not null comment('组织id') unique('roles_menus_unique_org_role_path') INT"`
	RoleName string `json:"roleName" xorm:"not null comment('角色名称') unique('roles_menus_unique_org_role_path') VARCHAR(32)"`
	MenuPath string `json:"menuPath" xorm:"not null comment('菜单路径') unique('roles_menus_unique_org_role_path') index CHAR(10)"`
}

func (m *RolesMenus) TableName() string {
	return "roles_menus"
}
