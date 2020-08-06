package admins

import (
	"zhuan-qian/go-saas/app/controllers"
	"zhuan-qian/go-saas/model"
	"zhuan-qian/go-saas/service"
)

type Locations struct {
	controllers.Base
}

//获取区域列表
func (c *Locations) Get() {
	var (
		m   []*model.Locations
		err error
	)

	m, err = service.NewLocationsService().List()
	if err != nil {
		c.SendServerError(err.Error())
		return
	}

	c.SendSmile(m)
}
