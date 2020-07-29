package service

import (
	"gold_hill/mine/common"
	"gold_hill/mine/dao"
	"gold_hill/mine/model"
	"strconv"
	"strings"
)

type RolesRbac struct {
}

func NewRolesRbacService() *RolesRbac {
	return &RolesRbac{}
}

func (s *RolesRbac) List(orgId int, roleId int) (list []*model.RbacPolicy) {
	var (
		rolesDao = dao.NewRolesDao().WithSession(nil)
		roleName string
		err      error
	)

	//TODO 更改管理员判断
	roleName, err = rolesDao.NameById(roleId)
	if err != nil {
		return nil
	}
	king, _ := rolesDao.IsKing(roleName)
	if king {
		return append(list, &model.RbacPolicy{
			RoleName: model.ROLES_NAME_OF_KING,
			Obj:      model.OBJ_OF_ADMIN,
			Act:      model.ACT_OF_ADMIN,
		})
	}

	return dao.NewRbacDao().ObjsByRoleName(orgId, roleName)
}

//编辑角色权限
func (s *RolesRbac) Modify(orgId int, roleName string, policies []*model.RbacPolicyWithoutRoleName) (bool, error) {
	var (
		rbac     = common.GetCasbin()
		err      error
		isKing   bool
		orgIdStr = strconv.Itoa(orgId)
	)
	isKing, err = dao.NewRolesDao().WithSession(nil).IsKing(roleName)
	if err != nil {
		return false, err
	}
	if isKing {
		return false, common.NewRequireError("管理员角色无法修改")
	}

	//rbac.RemovePolicies(rbac.GetPermissionsForUserInDomain(roleName, orgIdStr))
	for _, v := range rbac.GetPermissionsForUserInDomain(roleName, orgIdStr) {
		_, err = rbac.RemovePolicy(v[model.RbacRoleIndexOfPolicy], v[model.RbacDomainIndexOfPolicy],
			v[model.RbacObjIndexOfPolicy], v[model.RbacActIndexOfPolicy])
	}
	if err != nil {
		rbac.LoadPolicy()
		return false, err
	}
	for _, v := range policies {
		v.Obj = common.ParseUriToObj(v.Obj)
		_, err = rbac.AddPolicy(roleName, orgIdStr, strings.ToLower(v.Obj), strings.ToLower(v.Act))
		if err != nil {
			rbac.LoadPolicy()
			return false, err
		}
	}

	rbac.SavePolicy()
	return true, nil
}
