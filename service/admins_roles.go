package service

import (
	"gold_hill/mine/common"
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
	"strconv"
)

type AdminsRoles struct {
}

func NewAdminsRolesService() *AdminsRoles {
	return &AdminsRoles{}
}

func (s *AdminsRoles) ListBy(adminId int) (list []*model.Roles, err error) {
	var (
		rolesDao = dao.NewRolesDao().WithSession(nil)
	)

	list, err = rolesDao.ListByUserId(model.ORGANIZEID_OF_BACKEND, adminId)
	if err != nil {
		return nil, err
	}

	return
}

func (s *AdminsRoles) Modify(adminId int64, rolesId []string, organizeId int) (result int, err error) {
	var (
		rbac           = common.GetCasbin()
		adminIdStr     = strconv.FormatInt(adminId, 10)
		existRoles     []string
		validRoles     []*model.Roles
		validRolesName []string
		rolesDao       = dao.NewRolesDao().WithSession(nil)
		orgStr         = strconv.Itoa(organizeId)
	)
	defer rbac.SavePolicy()

	existRoles = rbac.GetRolesForUserInDomain(adminIdStr, orgStr)
	for _, role := range existRoles {
		rbac.DeleteRoleForUserInDomain(adminIdStr, role, orgStr)
	}

	if len(rolesId) < 1 {
		return 0, nil
	}

	validRoles, err = rolesDao.ListByIds([]string{"id", "name"}, organizeId, rolesId)
	if err != nil {
		return
	}
	validRolesName = rolesDao.CollectNamesBy(validRoles)

	if len(validRolesName) < 1 {
		return
	}
	for _, v := range validRolesName {
		rbac.AddRoleForUserInDomain(adminIdStr, v, orgStr)
	}

	return len(validRolesName), nil
}
