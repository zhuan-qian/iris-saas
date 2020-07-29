package dao

import (
	"gold_hill/mine/model"
	"gold_hill/mine/model/mashup"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type OrgGroups struct {
	Base
}

func NewOrgGroupsDao() *OrgGroups {
	return &OrgGroups{}
}

func (d *OrgGroups) WithSession(s *xorm.Session) *OrgGroups {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

//获取列表
func (d *OrgGroups) GetAll(limit *int, offset *int) (m []*model.OrgGroups, count int64, err error) {
	session := d.session.Cols("id", "name")

	cond := builder.NewCond()
	cond = cond.And(builder.Eq{"status": 1})
	if cond != builder.NewCond() {
		session = session.Where(cond)
	}

	if *limit != 0 {
		session = session.Limit(*limit, *offset)
	}
	count, err = session.FindAndCount(&m)
	return
}

func (d *OrgGroups) CollectIdBy(list []*model.OrgGroups) (ids []int) {
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	return
}

func (d *OrgGroups) CreateOne(u *model.OrgGroups) (int64, error) {
	return d.InsertOne(u)
}

func (d *OrgGroups) FillCountBy(list []*model.OrgGroups, mapper mashup.OrgWithGroupMap) {
	for i, v := range list {
		if _, ok := mapper[v.Id]; ok {
			list[i].Count = mapper[v.Id].Count
		}
	}
}
