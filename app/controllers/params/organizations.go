package params

type OrganizationsGet struct {
	Keyword  *string `json:"keyword,omitempty" validate:"omitempty,gte=1"`
	GroupId  *int    `json:"groupId,omitempty" validate:"omitempty,numeric"`
	Limit    *int    `json:"limit,omitempty" validate:"omitempty,numeric,gte=10,lte=60"`
	Page     *int    `json:"page,omitempty" validate:"omitempty,numeric,gte=1"`
	Sort     *string `json:"sort,omitempty" validate:"omitempty"`
	SortType *string `json:"sortType,omitempty" validate:"omitempty"`
}

type OrganizationsPut struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,gte=2,lte=24"`
	Phone *string `json:"phone,omitempty" validate:"omitempty,len=11"`
	//ExpireAt *string `json:"expire_at,omitempty" validate:"omitempty"`
	GroupId *int    `json:"group_id,omitempty" validate:"omitempty,numeric"`
	Area    *string `json:"area,omitempty" validate:"omitempty"`
	Address *string `json:"address,omitempty" validate:"omitempty,lte=200"`
	Status  *int8   `json:"status,omitempty" validate:"omitempty,numeric,oneof=-1 0 1"`
}

type OrganizationsPost struct {
	Owner *int64  `json:"owner" validate:"required,numeric"`
	Name  *string `json:"name" validate:"required,gte=2,lte=24"`
	Phone *string `json:"phone" validate:"required,len=11"`
	//Password *string `json:"password" validate:"required,gte=6,lte=32"`
	GroupId *int    `json:"group_id,omitempty" validate:"omitempty,numeric"`
	Area    *string `json:"area" validate:"required"`
	Address *string `json:"address" validate:"required,lte=200"`
	Type    *int8   `json:"type" validate:"required,numeric,oneof=0 1"`
	Status  *int8   `json:"status,omitempty" validate:"omitempty,numeric,oneof=-1 0 1"`
}
