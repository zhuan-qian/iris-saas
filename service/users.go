package service

import (
	"errors"
	"github.com/iris-contrib/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"os"
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/common"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service/cache"
	"zhuan-qian/go-saas/service/sms"
	"strconv"
	"strings"
	"time"
)

type users struct {
	d *dao.Users
}

func NewUsersService() *users {
	return &users{d: dao.NewUsersDao().WithSession(nil)}
}

//登录模块
func (s *users) LoginByPassword(account string, requestPassword []byte) (string, error) {
	var (
		token string
		m     *model.Users
		err   error
	)

	m, err = s.d.GetByAccount(account)
	if err != nil {
		return token, err
	}
	if m == nil {
		return token, errors.New("账号或密码错误")
	}
	passwordRecord := []byte(m.Password)
	err = bcrypt.CompareHashAndPassword(passwordRecord, requestPassword)
	if err != nil {
		return token, errors.New("账号或密码错误")
	}

	return s.CreateTokenAndCache(m)
}

func (s *users) LoginOrRegister(account string, code string) (token string, err error) {
	var (
		exist bool
	)

	exist, err = s.d.ExistBy(account)
	if err != nil {
		return "", err
	}
	if exist {
		token, err = s.LoginByCode(account, code)
	} else {
		token, err = s.CreateBy(account, code)
	}
	if err != nil {
		return "", err
	}
	_, err = s.d.UpdateLastLoginAtBy(account)
	return token, err
}

func (s *users) LoginByCode(account string, code string) (token string, err error) {
	var (
		m *model.Users
	)

	if !sms.NewAliSmsService().CodeIsRightThenDel(account, code) {
		return "", common.NewRequireError("账号或验证码错误")
	}

	m, err = s.d.GetByAccount(account)
	if err != nil {
		return token, err
	}

	token, err = s.CreateTokenAndCache(m)
	if err != nil {
		return token, err
	}

	err = s.SetCacheOfToken(token, m.Id)
	if err != nil {
		return token, errors.New("用户token存储失败 请联系管理员 " + err.Error())
	}

	return token, nil

}

//创建token
func (s *users) CreateTokenAndCache(m *model.Users) (string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        m.Id,
		"account":   m.Account,
		"nickname":  m.Nickname,
		"avatar":    m.Avatar,
		"pushToken": m.PushToken,
		"sex":       m.Sex,
		"orgId":     m.OrgId,
		"saleable":  m.Saleable,
		"createdAt": *m.CreatedAt,
		"status":    m.Status,
		"exp":       time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("USER_JWT_SECRET")))
	if err != nil {
		return "", err
	}
	err = s.SetCacheOfToken(tokenString, m.Id)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

//解析token存入控制器变量容器中(ctx.values())
func (s *users) ParseTokenToUser(jwt jwt.MapClaims) *model.Users {
	m := &model.Users{}
	m.Id = int64(jwt["id"].(float64))
	m.Account = jwt["account"].(string)
	m.Nickname = jwt["nickname"].(string)
	m.Avatar = jwt["avatar"].(string)
	m.Sex = byte(jwt["sex"].(float64))
	m.OrgId = int(jwt["orgId"].(float64))
	m.Saleable = int8(jwt["saleable"].(float64))
	m.Status = int8(jwt["status"].(float64))
	m.CreatedAt = common.StringPtr(jwt["createdAt"].(string))
	return m
}

//创建一个账号
func (s *users) Create(m *model.Users) (result int64, err error) {
	return s.d.CreateOne(m)
}

func (s *users) CreateBy(phone string, code string) (token string, err error) {
	var (
		m = &model.Users{
			Id:        0,
			Account:   phone,
			Password:  "",
			Nickname:  "",
			Avatar:    "",
			Sex:       2,
			PushToken: "",
			Status:    1,
			CreatedAt: common.NowDateTimePtr(),
		}
	)

	if !sms.NewAliSmsService().CodeIsRightThenDel(phone, code) {
		return "", common.NewRequireError("输入的验证码不正确")
	}
	rand.Seed(time.Now().UnixNano())
	m.Nickname = "用户" + strconv.Itoa(rand.Intn(9999999)) + m.Account[7:]

	_, err = s.Create(m)
	if err != nil {
		return "", err
	}
	return s.CreateTokenAndCache(m)
}

//设置token与用户id的映射缓存
func (s *users) SetCacheOfToken(token string, userId int64) error {
	return cache.RedisInit().HSet(string(cache.USERS_TOKEN_MAP), string(userId), token).Err()
}

//获取token相关的用户id
func (s *users) GetTokenByUid(userId int64) (string, error) {
	jwt, err := cache.RedisInit().HGet(string(cache.USERS_TOKEN_MAP), string(userId)).Result()
	if err != nil {
		return "0", err
	}
	return jwt, nil
}

//删除token相关的缓存
func (s *users) DelCacheOfToken(userId int64) error {
	return cache.RedisInit().HDel(string(cache.USERS_TOKEN_MAP), string(userId)).Err()
}

//通过关键词查询用户列表
func (s *users) GetListByKey(limit int, page int, keyword string) ([]*model.Users, int64, error) {
	keyword = strings.TrimSpace(keyword)
	return s.d.GetAll(limit, common.PageToOffset(page, limit), keyword, keyword)
}

//用户信息修改
func (s *users) Modify(id int64, p *params.UsersModify) (ok bool, err error) {
	var (
		m    = &model.Users{}
		cols []string
	)

	if p.Password != nil {
		if p.ConfirmPassword != nil {
			if p.Password == p.ConfirmPassword {
				err = s.DelCacheOfToken(id)
				if err != nil {
					return false, err
				}
				m.Password = s.d.EncryptPassword(*p.Password)
				cols = append(cols, "password")
			} else {
				return false, nil
			}
		} else {
			return false, nil
		}
	}
	if p.OrgId != nil {
		m.OrgId = *p.OrgId
		cols = append(cols, "orgId")
	}
	if p.Nickname != nil {
		m.Nickname = *p.Nickname
		cols = append(cols, "nickname")
	}
	if p.Avatar != nil {
		m.Avatar = *p.Avatar
		cols = append(cols, "avatar")
	}
	if p.BornAt != nil {
		m.BornAt = p.BornAt
		cols = append(cols, "bornAt")
	}
	if p.PushToken != nil {
		m.PushToken = *p.PushToken
		cols = append(cols, "pushToken")
	}

	if p.DynamicLock != nil {
		cols = append(cols, "dynamicLock")
		cols = append(cols, "lockTimeTill")
	}
	if p.LastLoginAt != nil {
		m.LastLoginAt = p.LastLoginAt
		cols = append(cols, "lastLoginAt")
	}
	if p.Status != nil {
		m.Status = *p.Status
		cols = append(cols, "status")
	}
	if p.Saleable != nil {
		m.Saleable = *p.Saleable
		cols = append(cols, "saleable")
	}
	if p.TalkStatus != nil {
		m.TalkStatus = *p.TalkStatus
		cols = append(cols, "talkStatus")
	}
	if p.Sex != nil {
		m.Sex = *p.Sex
		cols = append(cols, "sex")
	}
	_, err = s.d.Modify(id, cols, m)
	s.DelCacheOfToken(id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *users) InfoBy(id int64) (*model.Users, error) {
	var (
		m   = &model.Users{}
		ok  bool
		err error
	)
	ok, err = s.d.InfoBy(nil, id, m)
	if err != nil {
		return nil, err
	}
	if ok == false {
		return nil, nil
	}

	return m, err
}

func (d *users) MapInfoByIds(ids []int64) (list map[int64]*model.Users, err error) {
	return d.MapInfoByIds(ids)
}
