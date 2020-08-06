package service

import (
	"zhuan-qian/go-saas/app/controllers/params"
	"zhuan-qian/go-saas/dao"
	"zhuan-qian/go-saas/model"
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
