package admins

import (
	"gold_hill/mine/app/controllers"
	"gold_hill/mine/model"
	"gold_hill/mine/service"
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
