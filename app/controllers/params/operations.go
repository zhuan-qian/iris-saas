package params

type OperationsPost struct {
	Name   string `json:"name" validate:"required,min=1"`
	Params string `json:"params" validate:"required,min=1"`
}

type OperationsPut struct {
	Params string `json:"params" validate:"required,min=1"`
}
