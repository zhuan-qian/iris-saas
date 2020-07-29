package dao

import (
	"fmt"
	"gold_hill/mine/model"
	"strings"
	"xorm.io/xorm"
)

type EsLostIndex struct {
	Base
}

func NewEsLostIndexDao() *EsLostIndex {
	return &EsLostIndex{}
}

func (d *EsLostIndex) WithSession(s *xorm.Session) *EsLostIndex {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *EsLostIndex) Create(index string, objId int64) (int64, error) {
	return d.InsertOne(&model.EsLostIndex{
		Index: index,
		ObjId: objId,
	})
}

func (d *EsLostIndex) List(limit int) (list []*model.EsLostIndex, err error) {
	if limit > 0 {
		d.session = d.session.Limit(limit)
	}
	err = d.session.Find(&list)
	return list, err
}

func (d *EsLostIndex) Delete(index string, objIds []string) error {
	sql := fmt.Sprintf("delete from es_lost_index where `index`='%s' and objId in (%s)", index, strings.Join(objIds, ","))
	_, err := d.session.Exec(sql)
	return err
}
