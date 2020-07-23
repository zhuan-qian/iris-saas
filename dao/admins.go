package dao

import (
	"golang.org/x/crypto/bcrypt"
	"scaffold/common"
	"scaffold/model"
	"strconv"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	KEY_FOR_ADMIN_INFO = "admin"
)

type Admins struct {
	Base
}

func NewAdminsDao() *Admins {
	return &Admins{}
}

func (d *Admins) WithSession(s *xorm.Session) *Admins {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

func (d *Admins) GetByAccount(account string) (*model.Admins, error) {
	m := &model.Admins{}
	has, err := d.session.Cols("id", "account", "nickname", "token", "password", "status").
		And("account=?", account).And("status=?", 1).Get(m)
	if !has {
		return nil, err
	}
	return m, err
}

func (d *Admins) GetAll(limit int, offset int, keyword *string) (list []*model.Admins, count int64, err error) {
	if keyword != nil {
		d.session = d.session.Where(builder.NewCond().
			Or(builder.Like{"nickname", "%" + *keyword + "%"}, builder.Like{"account", "%" + *keyword + "%"}))
	}
	count, err = d.session.Cols("id", "account", "nickname", "status", "createdAt").And("status!=?", -1).Limit(limit, offset).FindAndCount(&list)
	return
}

func (d *Admins) CollectIdBy(list []*model.Admins) (ids []int) {
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	return
}

func (d *Admins) ListByUserIds(userIds []int) (roleNames map[int][]string) {
	var (
		rbac = common.GetCasbin()
	)
	list := make(map[int][]string)
	for _, v := range userIds {
		var roleName = rbac.GetRolesForUserInDomain(strconv.Itoa(v), strconv.Itoa(model.ORGANIZEID_OF_BACKEND))
		list[v] = roleName
	}

	return list
}

func (d *Admins) FillRoleNameBy(list []*model.Admins, roleNames map[int][]string) {
	for i, v := range list {
		if _, ok := roleNames[v.Id]; ok {
			list[i].RoleNames = roleNames[v.Id]
		}
	}
}

//加密密码
func (d *Admins) EncryptPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}

func (d *Admins) Modify(id int, cols []string, m *model.Admins) (int64, error) {
	return d.session.Cols(cols...).ID(id).Update(m)
}

func (d *Admins) Create(m *model.Admins) (int64, error) {
	return d.InsertOne(m)

}
