package admins

import (
	"gold_hill/scaffold/app/controllers"
	"gold_hill/scaffold/model"
	"gold_hill/scaffold/service"
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
