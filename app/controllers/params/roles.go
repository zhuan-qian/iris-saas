package params

type RolesGet struct {
	Keyword *string `json:"keyword,omitempty" validate:"omitempty,gte=1"`
	Limit   *int    `json:"limit,omitempty" validate:"omitempty,numeric,gte=10,lte=60"`
	Page    *int    `json:"page,omitempty" validate:"omitempty,numeric,gte=1"`
}

type RolesPut struct {
	OrgId  *int    `json:"orgId" validate:"omitempty,numeric,min=0"`
	Name   *string `json:"name" validate:"omitempty,gte=2,lte=24"`
	Status *int8   `json:"status" validate:"omitempty,oneof=-1 0 1"`
}
