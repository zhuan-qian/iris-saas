package dao

import (
	"scaffold/common"
	"scaffold/model"
	"strconv"
	"strings"
)

type Rbac struct {
}

func NewRbacDao() *Rbac {
	return &Rbac{}
}

func (s *Rbac) ObjsByRoleName(domain int, roleName string) (list []*model.RbacPolicy) {
	var (
		rbac  = common.GetCasbin()
		rbacs model.RbacPolicies
		maps  = make(map[string][]*model.RbacPolicy)
	)

	rbacs = rbac.GetPermissionsForUserInDomain(roleName, strconv.Itoa(domain))
	maps = s.PermissionsToRolesPolicies(&rbacs)
	return maps[roleName]
}

func (s *Rbac) ObjsByUserId(orgId int, userId int) (list []*model.RbacPolicy, err error) {
	var (
		rbac  = common.GetCasbin()
		rbacs model.RbacPolicies
		roles []string
	)

	roles = rbac.GetRolesForUserInDomain(strconv.Itoa(userId), strconv.Itoa(orgId))
	rbacs = s.CollectRolesPolicies(orgId, roles)
	return s.PermissionsToPolicies(&rbacs), nil
}

func (s *Rbac) CollectRolesPolicies(orgId int, roles []string) (rbacs model.RbacPolicies) {
	var (
		rbac = common.GetCasbin()
	)

	for _, role := range roles {
		rbacs = append(rbacs, rbac.GetPermissionsForUserInDomain(role, strconv.Itoa(orgId))...)
	}
	return rbacs
}

func (s *Rbac) PermissionsToRolesPolicies(permissions *model.RbacPolicies) map[string][]*model.RbacPolicy {
	var (
		maps = make(map[string][]*model.RbacPolicy)
	)

	for _, v := range *permissions {
		roleName := v[model.RbacRoleIndexOfPolicy]
		obj := v[model.RbacObjIndexOfPolicy]
		act := v[model.RbacActIndexOfPolicy]
		if _, ok := maps[roleName]; ok {
			maps[roleName] = append(maps[roleName], &model.RbacPolicy{
				RoleName: roleName,
				Obj:      obj,
				Act:      act,
			})
		} else {
			maps[roleName] = []*model.RbacPolicy{{
				RoleName: roleName,
				Obj:      obj,
				Act:      act,
			}}
		}
	}
	return maps
}

func (s *Rbac) PermissionsToPolicies(permissions *model.RbacPolicies) []*model.RbacPolicy {
	var (
		list []*model.RbacPolicy
	)

	for _, v := range *permissions {
		roleName := v[model.RbacRoleIndexOfPolicy]
		obj := v[model.RbacObjIndexOfPolicy]
		act := v[model.RbacActIndexOfPolicy]
		list = append(list, &model.RbacPolicy{
			RoleName: roleName,
			Obj:      obj,
			Act:      act,
		})
	}
	return list
}

func (s *Rbac) GetRolesNameBy(orgId int, userId int) []string {
	var (
		rbac = common.GetCasbin()
	)
	return rbac.GetRolesForUserInDomain(strconv.Itoa(userId), strconv.Itoa(orgId))
}

func (s *Rbac) ParseGroupPoliciesToRoles(g model.RbacGroupPolicies) (roles []string) {

	for _, v := range g {
		roles = append(roles, v[model.RbacRoleIndexOfGroupPolicy])
	}
	return
}

func (s *Rbac) ParsePoliciesToObjs(g model.RbacPolicies) (objs []string) {
	for _, v := range g {
		objs = append(objs, v[model.RbacObjIndexOfPolicy])
	}
	return
}

func (s *Rbac) ParsePoliciesToObjsAndActsMap(g model.RbacPolicies) (objs map[string]bool) {
	objs = make(map[string]bool)
	for _, v := range g {
		objs[v[model.RbacObjIndexOfPolicy]+","+v[model.RbacActIndexOfPolicy]] = true
	}
	return
}

func (s *Rbac) ValidRbacByUserId(userId int, organizeId int, uri string, act string) (result bool, err error) {
	var (
		rbac = common.GetCasbin()
		obj  string
	)

	//判断用户是否有接口操作权限
	obj = common.ParseUriToObj(uri)
	return rbac.Enforce(strconv.Itoa(userId), strconv.Itoa(organizeId), strings.ToLower(obj), strings.ToLower(act))
}
