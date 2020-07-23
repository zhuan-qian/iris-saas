package model

const (
	ORGANIZEID_OF_BACKEND = 0
)

type Organizations struct {
	Id        int     `json:"id" xorm:"not null pk autoincr INT"`
	Code      string  `json:"code" xorm:"not null comment('组织编码') VARCHAR(20) index"`
	Owner     int64   `json:"owner" xorm:"not null comment('所有人id') BIGINT index"`
	Name      string  `json:"name" xorm:"not null comment('店铺名') VARCHAR(32) index"`
	Phone     string  `json:"phone" xorm:"not null comment('手机') VARCHAR(16) index"`
	Area      string  `json:"area" xorm:"not null comment('区域') varchar(255) index"`
	Address   string  `json:"address" xorm:"not null comment('地理位置') VARCHAR(255)"`
	GroupId   int     `json:"groupId" xorm:"null comment('学校分组id') index INT"`
	Type      int8    `json:"type" xorm:"not null default 0 comment('组织类型 0: 多人 1: 个人') TINYINT"`
	Status    int8    `json:"status" xorm:"not null default 1 comment('状态 -1:删除 0:禁用 1:正常') TINYINT"`
	ExpireAt  *string `json:"expireAt,omitempty" xorm:"null comment('授权到期时间') DATE"`
	CreatedAt *string `json:"createdAt,omitempty" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt *string `json:"updatedAt,omitempty" xorm:"updated TIMESTAMP"`
}

func (m *Organizations) TableName() string {
	return "organizations"
}
