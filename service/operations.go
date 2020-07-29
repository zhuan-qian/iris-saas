package service

import (
	"gold_hill/mine/app/controllers/params"
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
)

type Operations struct {
	d *dao.Operations
}

func NewOperationsService() *Operations {
	return &Operations{d: dao.NewOperationsDao().WithSession(nil)}
}

func (s *Operations) Create(p *params.OperationsPost) error {
	m := &model.Operations{
		Name:   p.Name,
		Params: p.Params,
	}
	_, err := s.d.InsertOne(m)
	return err
}

func (s *Operations) ModifyBy(name string, p *params.OperationsPut) error {
	m := &model.Operations{
		Params: p.Params,
	}
	_, err := s.d.ModifyByField([]string{"params"}, "name", name, m)
	return err

}

func (s *Operations) DeleteBy(name string) error {
	err := s.d.DeleteByField("operations", "name", name)
	return err
}
