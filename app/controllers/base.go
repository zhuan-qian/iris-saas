package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"net/http"
	"gold_hill/mine/common"
)

type Base struct {
	Ctx      iris.Context
	Session  *sessions.Session
	Validate *validator.Validate
}

func CompactListAndCount(list interface{}, count int64) *map[string]interface{} {
	return &map[string]interface{}{"list": list, "count": count}
}

//func (c *Base) writeLog(){
//	c.Ctx.Application().Logger().Infof(c.Ctx.)
//}

func (c *Base) SendResponse(statusCode int, msg string, data interface{}) {
	c.Ctx.StatusCode(statusCode)
	box := make(map[string]interface{})
	box["msg"] = msg
	box["data"] = data
	c.Ctx.JSON(box)
	return
}

func (c *Base) SendCreated(data interface{}) {
	c.Ctx.StatusCode(http.StatusCreated)
	box := make(map[string]interface{})
	box["msg"] = "创建成功"
	box["data"] = data
	c.Ctx.JSON(box)
	return
}

func (c *Base) SendSmile(data interface{}) {
	c.Ctx.StatusCode(http.StatusOK)
	box := make(map[string]interface{})
	box["msg"] = "成功"
	box["data"] = data
	c.Ctx.JSON(box)
	return
}

func (c *Base) SendBadRequest(data interface{}, code *string) {
	c.Ctx.StatusCode(http.StatusBadRequest)
	box := make(map[string]interface{})
	box["msg"] = "请求失败"
	box["data"] = data
	box["code"] = code
	if box["code"] == nil {
		box["code"] = common.RESPONSE_EXPLAIN_BY_MSG
	}
	c.Ctx.JSON(box)
	return
}

func (c *Base) SendServerError(data interface{}) {
	c.Ctx.StatusCode(http.StatusInternalServerError)
	box := make(map[string]interface{})
	box["msg"] = "请求失败,内部错误"
	box["data"] = data
	c.Ctx.JSON(box)
	return
}
