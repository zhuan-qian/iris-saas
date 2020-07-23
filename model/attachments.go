package model

type Attachments struct {
	Id        int64   `json:"id" xorm:"not null pk autoincr BIGINT"`
	Mime      string  `json:"mime" xorm:"not null comment('文件类型') VARCHAR(50)"`
	Size      int64   `json:"size" xorm:"comment('文件尺寸') INT"`
	Path      string  `json:"path" xorm:"not null comment('路径') VARCHAR(2000)"`
	Sha1      string  `json:"sha1" xorm:"not null comment('文件sha1值') unique BINARY(40)"`
	Status    int8    `json:"status" xorm:"not null default 1 comment('文件状态 0:停用 1:正常') TINYINT"`
	CreatedAt *string `json:"createdAt,omitempty" xorm:"created not null default CURRENT_TIMESTAMP TIMESTAMP"`
	UpdatedAt *string `json:"updatedAt" xorm:"updated TIMESTAMP"`
}

func (m *Attachments) TableName() string {
	return "attachment"
}
