package client

import (
	"gold_hill/scaffold/app/controllers"
	"gold_hill/scaffold/service/sms"
)

type Sms struct {
	controllers.Base
}

type smsPost struct {
	Account string `json:"account" validate:"required,len=11"`
}

func (c *Sms) Post() {
	var (
		p   = &smsPost{}
		err = c.Ctx.ReadJSON(p)
	)

	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	err = c.Validate.Struct(*p)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	_, err = sms.NewTxSmsService().SendRandomMsg(p.Account)
	if err != nil {
		c.SendServerError(err.Error())
		return
	}
	c.SendSmile(true)
	return
}
