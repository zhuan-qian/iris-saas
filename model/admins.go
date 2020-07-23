package model

type Admins struct {
	Id        int     `json:"id" xorm:"not null pk autoincr INT"`
	Account   string  `json:"account" xorm:"not null comment('账号') unique(admins_unique_account) VARCHAR(16)"`
	Password  string  `json:"password,omitempty" xorm:"not null comment('密码') VARCHAR(255)"`
	Nickname  string  `json:"nickname" xorm:"not null comment('昵称') VARCHAR(16)"`
	Token     string  `json:"token,omitempty" xorm:"not null comment('令牌') VARCHAR(255) index"`
	Status    int8    `json:"status" xorm:"not null default 1 comment('状态 -1:删除 0:禁用 1:正常') TINYINT"`
	CreatedAt *string `json:"createdAt,omitempty" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt *string `json:"updatedAt" xorm:"updated TIMESTAMP"`

	RoleNames []string `json:"roleNames" xorm:"-"`
}

func (m *Admins) TableName() string {
	return "admins"
}
