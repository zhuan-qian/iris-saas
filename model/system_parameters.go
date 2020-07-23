package model

const (
	SYSTEM_PARAMETER_FOR_ADMIN_ROLE      = "role_for_admin_created" //管理员角色是否已创建
	SYSTEM_PARAMETER_FOR_ADMIN_ACCOUNT   = "admin_account_created"  //管理员账号是否已创建
	SYSTEM_PARAMETER_FOR_LOCATIONS_BUILT = "locations_built"        //全球城市信息是否已构建
	SYSTEM_PARAMETER_FOR_MENUS_BUILT     = "menus_built"            //菜单是否已构建
)

type SystemParameters struct {
	Name        string `json:"name" xorm:"not null comment('参数名') unique VARCHAR(64)"`
	Config      string `json:"config" xorm:"not null comment('配置') TEXT"`
	Description string `json:"description" xorm:"not null comment('描述') VARCHAR(255)"`
}

func (m *SystemParameters) TableName() string {
	return "system_parameters"
}
