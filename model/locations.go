package model

//全球地区库
type Locations struct {
	Id         int    `json:"id" xorm:"not null pk autoincr INT"`
	Path       string `json:"path" xorm:"not null comment('分类路径') VARCHAR(255) index"`
	Parent string `json:"parent" xorm:"not null comment('父级路径') VARCHAR(255) index"`
	Level      int    `json:"level" xorm:"not null comment('级别') INT"`
	Name       string `json:"name" xorm:"not null comment('名称') VARCHAR(64)"`
	NameEn     string `json:"nameEn" xorm:"not null comment('英文名称') VARCHAR(128)"`
	NamePinyin string `json:"namePinyin" xorm:"not null comment('拼音名称') VARCHAR(128)"`
	Code       string `json:"code" xorm:"not null comment('城市代码') VARCHAR(16)"`

	Subs []*Locations `json:"subs" xorm:"-"`
}

func (m *Locations) TableName() string {
	return "locations"
}
