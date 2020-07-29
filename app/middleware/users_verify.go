package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"net/http"
	"os"
	"gold_hill/mine/common"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
)

//用户token校验中间件
func UsersVerify(c iris.Context) {
	var (
		tokenStr    string
		cachedToken string
		user        *model.Users
		info        jwt.MapClaims
		token       *jwt.Token
		err         error
	)

	tokenStr = c.URLParam(common.TOKEN_NAME_IN_CLIENT)
	if tokenStr == "" {
		err = errors.New(common.TOKEN_NAME_IN_CLIENT + " JWT must in url")
		goto REQUEST_ERR
	}

	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("USER_JWT_SECRET")), nil
	})

	if err != nil {
		goto REQUEST_ERR
	}

	if token.Claims == nil {
		err = errors.New(common.TOKEN_NAME_IN_CLIENT + " header must be a valid JWT")
		goto REQUEST_ERR
	}

	info = token.Claims.(jwt.MapClaims)
	cachedToken, err = service.NewUsersService().GetTokenByUid(int64(info["id"].(float64)))
	if err != nil {
		err = errors.New("has an error : " + err.Error())
		goto SERVER_ERR
	}
	if tokenStr != cachedToken {
		err = errors.New(common.TOKEN_NAME_IN_CLIENT + "  expired or invalid")
		goto REQUEST_ERR
	}

	user = service.NewUsersService().ParseTokenToUser(info)
	c.Values().Set("user", user)
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
