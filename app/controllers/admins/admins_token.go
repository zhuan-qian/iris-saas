package admins

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/common"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
	"strings"
)

type AdminsToken struct {
	controllers.Base
}

//登录
func (c *AdminsToken) Post() {
	var (
		a     = &model.Admins{}
		token *string
		err   error
	)

	err = c.Ctx.ReadForm(a)
	if a.Account == "" || a.Password == "" {
		c.SendBadRequest("参数缺失", nil)
		return
	}

	token, err = service.NewAdminsService().Login(a.Account, []byte(a.Password))
	if err != nil {
		if common.IsRequireError(err) {
			c.SendBadRequest(err.Error(), nil)
		} else {
			c.SendServerError(err.Error())
		}
		return
	}
	c.SendSmile(token)
	return
}

//退出登录
func (c *AdminsToken) Delete() {
	var (
		token           *jwt.Token
		tokenStr        string
		admin           jwt.MapClaims
		authHeaderParts []string
		err             error
	)

	tokenStr = c.Ctx.GetHeader(common.TOKEN_NAME_IN_BACKEND)
	if tokenStr == "" {
		c.SendBadRequest("unauthorized of token", nil)
		return
	}

	authHeaderParts = strings.Split(tokenStr, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		c.SendBadRequest("token header format must be Bearer {token}", nil)
		return
	}
	tokenStr = authHeaderParts[1]

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("ADMIN_TOKEN_SECRET")), nil
	})

	if err != nil {
		c.SendBadRequest("token failure "+err.Error(), nil)
		return
	}

	if token.Claims == nil {
		c.SendBadRequest("token failure", nil)
		return
	}

	admin = token.Claims.(jwt.MapClaims)
	cachedToken, err := service.NewAdminsService().GetCacheByUid(int64(admin["id"].(float64)))
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	if tokenStr != cachedToken {
		c.SendBadRequest("token Matching failure", nil)
		return
	}

	err = service.NewAdminsService().DelCacheOfToken(int64(admin["id"].(float64)))

	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(nil)
}
