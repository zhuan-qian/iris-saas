package client

import (
	"scaffold/app/controllers"
)

type AlipaySync struct {
	controllers.Base
}

func (c *AlipaySync) Post() {
	//var(
	//	err error
	//)

	//err=service.NewAlipayService().SyncResultForPay(c.Ctx.Request())
	//if err!=nil{
	//	c.SendServerError(err.Error())
	//	return
	//}
	//c.SendSmile("success")
	//return

//REQUEST_ERR:
//	c.SendBadRequest(err.Error(), common.StringPtr(common.RESPONSE_CRETE_ORDER_ERR))
//	return
//
//SERVICE_ERR:
//	c.SendServerError(err.Error())
//	return

}
