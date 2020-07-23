package model

const (
	DYNAMIC_LOCK_IS_OPEN      = 1
	DYNAMIC_LOCK_IS_CLOSE     = 0
	USERS_SALEABLE_IS_DISABLE = 0
	USERS_SALEABLE_IS_ENABLE  = 1
	USERS_STATUS_DELETE       = -1
	USERS_STATUS_DISABLE      = 0
	USERS_STATUS_NORMAL       = 1
)

//用户表
type Users struct {
	Id           int64   `json:"id" xorm:"not null pk autoincr BIGINT"`
	WxId         int     `json:"wxId" xorm:"null comment('微信id') INT"`
	Account      string  `json:"account" xorm:"not null comment('账号') unique(account) VARCHAR(16)" validate:"required,len=11"`
	Password     string  `json:"password,omitempty" xorm:"not null comment('密码') VARCHAR(255)" validate:"required,gte=6"`
	Nickname     string  `json:"nickname" xorm:"not null comment('昵称') VARCHAR(32) index" validate:"required,gte=2,lte=24"`
	Avatar       string  `json:"avatar" xorm:"not null comment('头像url') VARCHAR(255)" validate:"required,url"`
	Sex          byte    `json:"sex" xorm:"not null comment('性别 0:女 1:男 2:保密') TINYINT" validate:"required,numeric,oneof=0 1 2"`
	BornAt       *string `json:"bornAt" xorm:"null comment('出生时间') DATE" validate:"datetime=2006-01-02"`
	PushToken    string  `json:"pushToken,omitempty" xorm:"null comment('极光推送token') VARCHAR(128)"`
	NewbieGuided byte    `json:"newbieGuided" xorm:"not null comment('新手引导是否结束 0:否 1:是') TINYINT" validate:"numeric,oneof=0 1"`
	LastOnlineAt *string `json:"lastOnlineAt,omitempty" xorm:"comment('最后在线时间') TIMESTAMP"`
	LastLoginAt  *string `json:"lastLoginAt,omitempty" xorm:"comment('最后登录时间') TIMESTAMP"`
	OrgId        int     `json:"orgId,omitempty" xorm:"not null comment('店铺id') INT"`
	Saleable     int8    `json:"saleable" xorm:"not null default 0 comment('是否可开店销售 0:否 1:是') TINYINT"`
	Status       int8    `json:"status" xorm:"not null default 1 comment('状态 -1:删除 0:禁用 1:正常') TINYINT" validate:"numeric,oneof=-1 0 1"`
	TalkStatus   int8    `json:"talkStatus" xorm:"not null default 1 comment('状态 0:禁言 1:发言') TINYINT" validate:"numeric,oneof=0 1"`
	CreatedAt    *string `json:"createdAt,omitempty" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt    *string `json:"updatedAt,omitempty" xorm:"updated TIMESTAMP"`
}

func (m *Users) TableName() string {
	return "users"
}
