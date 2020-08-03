package model

//全球地区库
type Locations struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT"`
	Path       string `json:"path" xorm:"not null comment('分类路径') VARCHAR(255) index"`
	Parent     string `json:"parent" xorm:"not null comment('父级路径') VARCHAR(255) index"`
	Level      int8   `json:"level" xorm:"not null comment('级别') TINYINT"`
	Name       string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	Initial     string `json:"initial" xorm:"not null comment('拼音名称') CHAR(1)"`
	NamePinyin string `json:"namePinyin" xorm:"not null comment('拼音名称') VARCHAR(128)"`
	Subs []*Locations `json:"subs" xorm:"-"`
}

func (m *Locations) TableName() string {
	return "locations"
}
