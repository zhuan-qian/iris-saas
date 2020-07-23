package params

type RolesMenusPost struct {
	MenusPath string `json:"menusPath" validate:"required,gte=8"`
}
