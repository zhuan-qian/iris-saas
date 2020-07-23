package params

type OrgGroupsGet struct {
	Limit *int `json:"limit,omitempty" validate:"omitempty,numeric,gte=10,lte=60"`
	Page  *int `json:"page,omitempty" validate:"omitempty,numeric,gte=1"`
}

type OrgGroupsPut struct {
	Name   *string `json:"name,omitempty" validate:"omitempty,gte=2,lte=24"`
	Status *int8   `json:"status,omitempty" validate:"omitempty,numeric,oneof=-1 1"`
}
