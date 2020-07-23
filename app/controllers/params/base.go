package params

type Base struct {
	JWT *string `json:"JWT,omitempty" uri:"JWT" validate"omitempty"`
}
