package dao

import (
	"fmt"
	"gold_hill/mine/model"
	"gold_hill/mine/service/cache"
	"strings"
	"time"
	"xorm.io/xorm"
)

type EsLostUpdate struct {
	Base
}

func NewEsLostUpdateDao() *EsLostUpdate {
	return &EsLostUpdate{}
}

func (d *EsLostUpdate) WithSession(session *xorm.Session) *EsLostUpdate {
	if session != nil {
		d.Write(session)
	} else {
		d.NewSession()
	}
	return d
}

func (d *EsLostUpdate) Create(index string, objId int64) (int64, error) {
	success, err := cache.RedisInit().SetNX(d.cacheKey(index, objId), "1", time.Minute*30).Result()
	if err != nil {
		return 0, err
	}
	if success {
		return d.session.InsertOne(&model.EsLostUpdate{
			Index: index,
			ObjId: objId,
		})
	}
	return 0, nil
}

func (d *EsLostUpdate) cacheKey(index string, objId int64) string {
	return fmt.Sprintf("%d_%s_%d", cache.ES_LOST_UPDATE, index, objId)
}

func (d *EsLostUpdate) List(limit int) (list []*model.EsLostUpdate, err error) {
	if limit > 0 {
		d.session = d.session.Limit(limit)
	}
	err = d.session.Find(&list)
	return list, err
}

func (d *EsLostUpdate) Delete(index string, objIds []string) error {
	sql := fmt.Sprintf("delete from es_lost_update where `index`='%s' and objId in (%s)", index, strings.Join(objIds, ","))
	_, err := d.session.Exec(sql)
	return err
}
