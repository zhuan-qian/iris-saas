package model

const (
	OBJ_OF_ADMIN                 = "all"
	ACT_OF_ADMIN                 = "all"
	RbacUserIndexOfGroupPolicy   = 0
	RbacRoleIndexOfGroupPolicy   = 1
	RbacDomainIndexOfGroupPolicy = 2
	RbacRoleIndexOfPolicy        = 0
	RbacDomainIndexOfPolicy      = 1
	RbacObjIndexOfPolicy         = 2
	RbacActIndexOfPolicy         = 3
)

type RbacGroupPolicies [][]string
type RbacPolicies [][]string

type RbacPolicy struct {
	RoleName string `json:"roleName" validate:"required,min=1"`
	Obj      string `json:"obj" validate:"required,startswith=/"`
	Act      string `json:"act" validate:"required,oneof=get post put delete"`
}

type RbacPolicyWithoutRoleName struct {
	Obj string `json:"obj" validate:"required,startswith=/"`
	Act string `json:"act" validate:"required,oneof=get post put delete"`
}

type RbacPolicyWithDescription struct {
	Obj         string `json:"obj" validate:"required,startswith=/"`
	Act         string `json:"act" validate:"required,oneof=get post put delete"`
	Description string `json:"description"`
}
