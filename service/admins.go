package service

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service/cache"
	"strings"
	"time"
)

type Admins struct {
	d *dao.Admins
}

func NewAdminsService() *Admins {
	return &Admins{d: dao.NewAdminsDao().WithSession(nil)}
}

//登录模块
func (s *Admins) Login(account string, requestPassword []byte) (token *string, err error) {
	var (
		m *model.Admins
	)

	m, err = s.d.GetByAccount(account)
	if err != nil {
		return token, err
	}
	if m == nil {
		return token, common.NewRequireError("账号或密码错误")
	}
	passwordRecord := []byte(m.Password)
	err = bcrypt.CompareHashAndPassword(passwordRecord, requestPassword)
	if err != nil {
		return token, common.NewRequireError("账号或密码错误")
	}

	token, err = s.CreateToken(m)
	if err != nil {
		return token, err
	}

	err = s.SetCacheOfToken(*token, int64(m.Id))
	if err != nil {
		return token, errors.New("用户token存储失败 请联系管理员 " + err.Error())
	}

	return token, nil
}

//创建token
func (s *Admins) CreateToken(m *model.Admins) (*string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       m.Id,
		"account":  m.Account,
		"nickname": m.Nickname,
		"status":   m.Status,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("ADMIN_TOKEN_SECRET")))
	return &tokenString, err
}

//解析token存入控制器变量容器中(ctx.values())
func (s *Admins) ParseTokenToUser(jwt jwt.MapClaims) *model.Admins {
	m := &model.Admins{}
	m.Id = int(jwt["id"].(float64))
	m.Account = jwt["account"].(string)
	m.Nickname = jwt["nickname"].(string)
	m.Status = int8(jwt["status"].(float64))
	return m
}

//创建一个账号
func (s *Admins) Create(account string, password string, nickname string) (id int, err error) {
	m := &model.Admins{}
	m.Account, m.Password, m.Nickname, m.Status = account, s.d.EncryptPassword(password), nickname, 1
	_, err = s.d.Create(m)
	return m.Id, err
}

//设置token与管理员id的映射缓存
func (s *Admins) SetCacheOfToken(token string, userId int64) error {
	return cache.RedisInit().HSet(string(cache.ADMINS_TOKEN_MAP), string(userId), token).Err()
}

//获取token相关的管理员id
func (s *Admins) GetCacheByUid(userId int64) (string, error) {
	result, err := cache.RedisInit().HGet(string(cache.ADMINS_TOKEN_MAP), string(userId)).Result()
	if err != nil {
		return "0", err
	}
	return result, nil
}

//删除token相关的缓存
func (s *Admins) DelCacheOfToken(userId int64) error {
	return cache.RedisInit().HDel(string(cache.ADMINS_TOKEN_MAP), string(userId)).Err()
}

func (s *Admins) GetListByKey(limit int, page int, keyword *string) (list []*model.Admins, count int64, err error) {
	var (
		ids       []int
		roleNames map[int][]string
	)

	if keyword != nil {
		*keyword = strings.TrimSpace(*keyword)
	}
	list, count, err = s.d.GetAll(limit, common.PageToOffset(page, limit), keyword)
	if err != nil {
		return nil, 0, err
	}
	ids = s.d.CollectIdBy(list)
	roleNames = s.d.ListByUserIds(ids)
	s.d.FillRoleNameBy(list, roleNames)

	return
}

//员工信息修改
func (s *Admins) Modify(id int, p *params.AdminsModify) (ok bool, err error) {
	var (
		m    = &model.Admins{}
		cols []string
	)

	if p.Password != nil {
		s.DelCacheOfToken(int64(id))
		m.Password = s.d.EncryptPassword(*p.Password)
		cols = append(cols, "password")
	}
	if p.Nickname != nil {
		m.Nickname = *p.Nickname
		cols = append(cols, "nickname")
	}
	if p.Account != nil {
		s.DelCacheOfToken(int64(id))
		m.Account = *p.Account
		cols = append(cols, "account")
	}
	if p.Status != nil {
		m.Status = *p.Status
		cols = append(cols, "status")
	}
	_, err = s.d.Modify(id, cols, m)
	s.DelCacheOfToken(int64(id))
	if err != nil {
		return false, err
	}
	return true, nil
}
