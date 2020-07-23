package model

type OrgGroups struct {
	Id     int    `json:"id" xorm:"not null pk autoincr INT"`
	Name   string `json:"name" xorm:"not null comment('组名') unique(name) VARCHAR(32)" validate:"required,gte=2,lte=24"`
	Status int8   `json:"status" xorm:"not null default 1 comment('状态 -1:删除 1:正常') TINYINT" validate:"numeric,oneof=-1 1"`

	Count int `json:"count" xorm:"-"`
}

func (m *OrgGroups) TableName() string {
	return "org_groups"
}
