package dao

import (
	"gold_hill/mine/common"
	"gold_hill/mine/model"
	"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Roles struct {
	Base
}

func NewRolesDao() *Roles {
	return &Roles{}
}

func (d *Roles) WithSession(s *xorm.Session) *Roles {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Roles) ListByIds(cols []string, orgId int, rolesId []string) (roles []*model.Roles, err error) {
	err = d.session.Cols(cols...).Where("orgId=?", orgId).In("id", rolesId).Find(&roles)
	return
}

func (d *Roles) ListByUserId(orgId int, userId int) (roles []*model.Roles, err error) {
	var (
		rolesName []string
		rbac      = common.GetCasbin()
	)
	rolesName = rbac.GetRolesForUserInDomain(strconv.Itoa(userId), strconv.Itoa(orgId))
	err = d.session.Cols("id", "orgId", "name", "status", "createdAt").Where("orgId=?", orgId).
		In("name", rolesName).Find(&roles)
	return
}

func (d *Roles) CollectNamesBy(roles []*model.Roles) (rolesName []string) {
	for _, v := range roles {
		rolesName = append(rolesName, v.Name)
	}
	return
}

func (d *Roles) GetAll(orgId int, limit *int, offset *int, name *string) (m []model.Roles, count int64, err error) {
	session := d.session.Cols("id", "orgId", "name", "status", "createdAt").Where("orgId=?", orgId)

	cond := builder.NewCond()
	if *name != "" {
		cond = cond.Or(builder.Like{"name", *name + "%"})
	}
	if cond != builder.NewCond() {
		session = session.Where(cond)
	}

	if *limit != 0 {
		session = session.Limit(*limit, *offset)
	}
	count, err = session.FindAndCount(&m)
	return
}

func (d *Roles) IsKing(roleName string) (bool, error) {
	if roleName != model.ROLES_NAME_OF_KING {
		return false, nil
	}
	return true, nil
}

func (d *Roles) IsKingBy(orgId int, roleId int) (bool, error) {
	var (
		m = &model.Roles{}
	)
	return d.session.ID(roleId).Where("orgId=?", orgId).And("name=?", "admin").And("status=?", "1").Exist(m)
}

func (d *Roles) ContainsKing(rolesName []string) bool {
	for _, v := range rolesName {
		if v == model.ROLES_NAME_OF_KING {
			return true
		}
	}
	return false
}

func (d *Roles) NameById(roleId int) (name string, err error) {
	r := &model.Roles{}
	_, err = d.InfoBy([]string{"name"}, int64(roleId), r)
	if err != nil {
		return "", err
	}
	return r.Name, nil
}
