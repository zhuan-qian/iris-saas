package params

type MenusGet struct {
	IsTree      int8    `json:"isTree" validate:"required,numeric,oneof=0 1"`
	RolesName   *string `json:"rolesName,omitempty" validate:"omitempty,gte=1"`
	OnlyRelated *int8   `json:"onlyRelated,omitempty" validate:"omitempty,oneof=0 1"`
	TagRelated  *int8   `json:"tagRelated,omitempty" validate:"omitempty,oneof=0 1"`
}
