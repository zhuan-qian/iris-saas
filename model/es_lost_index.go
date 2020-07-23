package model

type EsLostIndex struct {
	Id    int    `json:"id" xorm:"not null pk autoincr INT"`
	Index string `json:"index" xorm:"not null comment('索引名称或别名') unique('index_obj_id') VARCHAR(255)"`
	ObjId int64  `json:"objId" xorm:"not null comment('对象id') unique('index_obj_id') BIGINT"`
}

func (m *EsLostIndex) TableName() string {
	return "es_lost_index"
}
