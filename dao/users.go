package dao

import (
	"golang.org/x/crypto/bcrypt"
	"gold_hill/mine/common"
	"gold_hill/mine/model"
	"xorm.io/builder"
	"xorm.io/xorm"
)

const (
	KEY_FOR_USER_INFO = "user"
)

type Users struct {
	Base
}

func NewUsersDao() *Users {
	return &Users{}
}

func (d *Users) WithSession(s *xorm.Session) *Users {
	if s != nil {
		d.Write(s)
	} else {
		d.NewSession()
	}
	return d
}

//加密密码
func (d *Users) EncryptPassword(password string) string {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(p)
}

//通过账号获取用户信息
func (d *Users) GetByAccount(account string) (*model.Users, error) {
	m := &model.Users{}
	has, err := d.session.And("account=?", account).And("status=?", 1).Get(m)
	if !has {
		return nil, err
	}
	return m, err
}

//用户信息获取
func (d *Users) Info(userId int64) (*model.Users, error) {
	user := new(model.Users)
	find, err := d.session.Omit("password", "token", "pushToken").ID(userId).Get(user)
	if err != nil {
		return nil, err
	}
	if !find {
		return nil, nil
	}
	return user, nil
}

func (d *Users) ExistBy(account string) (r bool, err error) {
	user := &model.Users{}
	r, _ = d.session.Where("account=?", account).Exist(user)
	return
}

//创建一个用户
func (d *Users) CreateOne(u *model.Users) (int64, error) {
	u.Password = d.EncryptPassword(u.Password)
	return d.InsertOne(u)
}

//获取用户列表
func (d *Users) GetAll(limit int, offset int, account string, nickname string) (m []*model.Users, count int64, err error) {
	session := d.session.Cols("id", "avatar", "account", "nickname", "sex", "bornAt", "createdAt", "status", "talkStatus").Where("status!=?", -1)

	cond := builder.NewCond()
	if account != "" {
		cond = cond.Or(builder.Like{"account", account + "%"})
	}
	if nickname != "" {
		cond = cond.Or(builder.Like{"nickname", nickname + "%"})
	}
	if cond != builder.NewCond() {
		session = session.Where(cond)
	}

	if limit != 0 {
		session = session.Limit(limit, offset)
	}
	count, err = session.FindAndCount(&m)
	return
}

//用户存在判断
func (d *Users) Exist(id int64) (bool, error) {
	m := &model.Users{}
	return d.session.ID(id).Exist(m)
}

//修改
func (d *Users) Modify(id int64, cols []string, m *model.Users) (int64, error) {
	return d.session.Cols(cols...).ID(id).Update(m)
}

//获取拥有token的用户
func (d *Users) HasPushTokenUsersBy(ids []int64) (list []model.Users, err error) {
	err = d.session.Cols("id", "pushToken").In("id", ids).Where("pushToken!=?", "").Find(&list)
	return
}

func (d *Users) UserPushTokenMapBy(hasPushTokenUsers []model.Users) *map[int64]string {
	var m = make(map[int64]string)

	for _, v := range hasPushTokenUsers {
		m[v.Id] = v.PushToken
	}
	return &m
}

func (d *Users) UpdateLastLoginAtBy(account string) (int64, error) {
	m := &model.Users{}
	now := common.NowDateTime()
	m.LastLoginAt = &now
	return d.session.Cols("lastLoginAt").Where("account=?", account).Update(m)
}

//根据用户id列表获取批量用户用户信息
func (d *Users) ListInfoByIds(ids []int64) (map[int64]*model.Users, error) {
	var (
		m   = make(map[int64]*model.Users)
		err error
	)
	err = d.session.In("id", ids).Where("status=1").Find(&m)
	return m, err
}

func (d *Users) ListBy(cols []string, ids []string) (list []*model.Users, err error) {
	if cols != nil {
		d.session = d.session.Cols(cols...)
	}
	err = d.session.In("id", ids).Where("status=1").Find(&list)
	return
}

func (d *Users) MapInfoByIds(ids []int64) (list map[int64]*model.Users, err error) {
	list = make(map[int64]*model.Users)
	err = d.session.In("id", ids).Where("status=1").Find(&list)
	return
}
