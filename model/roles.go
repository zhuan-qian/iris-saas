package model

const (
	ROLES_NAME_OF_KING = "admin"
)

type Roles struct {
	Id        int     `json:"id" xorm:"not null pk autoincr INT"`
	OrgId     int     `json:"orgId" xorm:"not null default 0 comment('组织id') unique('roles_unique_orgId_name') INT" validate:"numeric,min=0"`
	Name      string  `json:"name" xorm:"not null comment('角色名称') unique('roles_unique_orgId_name') VARCHAR(32)" validate:"required,gte=2,lte=24"`
	Status    int8    `json:"status" xorm:"not null default 1 comment('状态 0:禁用 1:正常') TINYINT" validate:"numeric,oneof=0 1"`
	CreatedAt *string `json:"createdAt,omitempty" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt *string `json:"updatedAt" xorm:"updated TIMESTAMP"`
}

func (m *Roles) TableName() string {
	return "roles"
}
