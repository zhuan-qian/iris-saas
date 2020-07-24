package service

import (
	"gold_hill/scaffold/dao"
	"gold_hill/scaffold/model"
)

type Rbac struct {
}

func NewRbacService() *Rbac {
	return &Rbac{}
}

func (s *Rbac) PoliciesBy(genre int8, menuPath *string) (err error, list []*model.MenusResources) {
	return dao.NewMenusResourcesDao().WithSession(nil).ListBy(genre, menuPath)
}
