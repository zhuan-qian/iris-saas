package service

import (
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
)

type AdminsRbac struct {
}

func NewAdminsRbacService() *AdminsRbac {
	return &AdminsRbac{}
}

func (s *AdminsRbac) List(adminId int) (list []*model.RbacPolicy, err error) {
	var (
		rolesName []string
		isKing    bool
	)

	rolesName = dao.NewRbacDao().GetRolesNameBy(model.ORGANIZEID_OF_BACKEND, adminId)
	isKing = dao.NewRolesDao().ContainsKing(rolesName)
	if isKing {
		return append(list, &model.RbacPolicy{
			RoleName: model.ROLES_NAME_OF_KING,
			Obj:      model.OBJ_OF_ADMIN,
			Act:      model.ACT_OF_ADMIN,
		}), nil
	}

	return dao.NewRbacDao().ObjsByUserId(model.ORGANIZEID_OF_BACKEND, adminId)
}
