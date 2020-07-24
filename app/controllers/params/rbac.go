package params

import "gold_hill/scaffold/model"

type RbacGet struct {
	MenuPath *string `json:"menuPath" validate:"omitempty,len=8"`
}

type RolesRbacPut struct {
	RoleName   string                             `json:"roleName" validate:"required,min=1"`
	PolicyList []*model.RbacPolicyWithoutRoleName `json:"policyList" validate:"required,dive,gte=1"`
}
