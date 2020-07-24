package dao

import (
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/model/mashup"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Organizations struct {
	Base
}

func NewOrganizationsDao() *Organizations {
	return &Organizations{}
}

func (d *Organizations) WithSession(s *xorm.Session) *Organizations {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Organizations) ListByids(cols []string, ids []int) (list []*model.Organizations, err error) {
	err = d.session.Cols(cols...).In("id", ids).Find(&list)
	return
}

//获取列表
func (d *Organizations) GetAll(limit *int, offset *int, name *string, phone *string, groupId *int, sort *string, sortType *string) (m []*model.Organizations, count int64, err error) {
	var (
		orderBy = "createdAt desc"
	)
	if *sortType != "" && *sort != "" {
		orderBy = *sortType + " " + *sort
	}
	session := d.session.Cols("id", "owner", "name", "groupId", "type", "phone", "area", "address", "status", "expireAt")

	cond := builder.NewCond()
	if *name != "" {
		cond = cond.Or(builder.Like{"name", *name + "%"})
	}
	if *phone != "" {
		cond = cond.Or(builder.Like{"phone", *phone + "%"})
	}
	if *groupId != 0 {
		cond = cond.And(builder.Eq{"groupId": *groupId})
	}
	if cond != builder.NewCond() {
		session = session.Where(cond)
	}

	if *limit != 0 {
		session = session.OrderBy(orderBy).Limit(*limit, *offset)
	}
	count, err = session.FindAndCount(&m)
	return
}

func (d *Organizations) CreateOne(m *model.Organizations) (int64, error) {
	return d.InsertOne(m)
}

func (d *Organizations) CountBy(groupIds []int) (list []*mashup.OrgWithGroup, err error) {
	err = d.session.Select("count('*') as count,groupId").In("groupId", groupIds).GroupBy("groupId").Find(&list)
	return
}

func (d *Organizations) MapOfGroupIdBy(list []*mashup.OrgWithGroup) mashup.OrgWithGroupMap {
	var (
		mapper = make(mashup.OrgWithGroupMap)
	)

	for _, v := range list {
		mapper[v.GroupId] = v
	}
	return mapper
}
