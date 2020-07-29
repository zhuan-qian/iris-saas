package middleware

import (
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
)

func AdminRbacVerify(c iris.Context) {
	var (
		valid bool
		err   error
		admin = c.Values().Get(dao.KEY_FOR_ADMIN_INFO).(*model.Admins)
	)

	valid, err = dao.NewRbacDao().ValidRbacByUserId(admin.Id, model.ORGANIZEID_OF_BACKEND, c.Path(), c.Method())
	if err != nil {
		goto SERVER_ERR
	}

	if !valid {
		err = errors.New("You don't have the permission of the resource! ")
		goto FORBIDDEN_ERR
	}

	c.Next()
	return

FORBIDDEN_ERR:
	c.StatusCode(http.StatusForbidden)
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
