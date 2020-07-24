package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
	"gold_hill/scaffold/common"
	"gold_hill/scaffold/dao"
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/service"
	"strings"
)

//管理员token校验中间件
func AdminsVerify(c iris.Context) {
	var (
		tokenStr        string
		cachedToken     string
		admin           *model.Admins
		token           *jwt.Token
		authHeaderParts []string
		info            jwt.MapClaims
		err             error
	)

	tokenStr = c.GetHeader(common.TOKEN_NAME_IN_BACKEND)
	if tokenStr == "" {
		err = errors.New(common.TOKEN_NAME_IN_BACKEND + " must in request header")
		goto REQUEST_ERR
	}

	authHeaderParts = strings.Split(tokenStr, " ")
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
		err = errors.New(common.TOKEN_NAME_IN_BACKEND + " header format must like as 'Bearer {token}'")
		goto REQUEST_ERR
	}
	tokenStr = authHeaderParts[1]

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("ADMIN_TOKEN_SECRET")), nil
	})

	if err != nil {
		goto REQUEST_ERR
	}

	if token.Claims == nil {
		err = errors.New(common.TOKEN_NAME_IN_BACKEND + " header must be a valid JWT")
		goto REQUEST_ERR
	}
	info = token.Claims.(jwt.MapClaims)
	cachedToken, err = service.NewAdminsService().GetCacheByUid(int64(info["id"].(float64)))
	if err != nil {
		goto SERVER_ERR
	}
	if tokenStr != cachedToken {
		err = errors.New(common.TOKEN_NAME_IN_BACKEND + " is expired or invalid")
		goto REQUEST_ERR
	}

	admin = service.NewAdminsService().ParseTokenToUser(info)

	c.Values().Set(dao.KEY_FOR_ADMIN_INFO, admin)
	c.Next()
	return

REQUEST_ERR:
	c.StatusCode(http.StatusUnauthorized)
	c.Header("WWW-Authenticate", fmt.Sprintf(`JWT realm="%s", charset="UTF-8"`, err.Error()))
	c.StopExecution()
	c.EndRequest()
	return

SERVER_ERR:
	c.StatusCode(http.StatusInternalServerError)
	c.Header("WWW-Authenticate", fmt.Sprintf(`JWT realm="%s", charset="UTF-8"`, err.Error()))
	c.StopExecution()
	c.EndRequest()
	return
}
