package service

import (
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
)

type Locations struct {
	d *dao.Locations
}

func NewLocationsService() *Locations {
	return &Locations{d: dao.NewLocationsDao().WithSession(nil)}
}

//查询中国区域列表
func (s *Locations) List() (list []*model.Locations, err error) {
	var (
		allList []*model.Locations
	)
	allList, err = s.d.ListByPath(",1,7,")
	if err != nil {
		return
	}
	list = s.d.BuildByTree(allList)
	s.d.Sort(true, list)
	return list, err
}
