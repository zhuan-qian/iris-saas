package params

type AdminsGet struct {
	Keyword *string `json:"keyword" validate:"omitempty,min=1"`
	Limit   int     `json:"limit" validate:"required,numeric,gte=10,lte=1000"`
	Page    int     `json:"page" validate:"required,numeric,min=1"`
}

type AdminsPost struct {
	Account  string `json:"account" validate:"required,numeric,len=11"`
	Password string `json:"password,omitempty" validate:"required,gte=6,lte=32"`
	Nickname string `json:"nickname" validate:"required,gte=2,lte=32"`
	Status   int8   `json:"status" validate:"required"`
}

type AdminsModify struct {
	Account  *string `json:"account" validate:"omitempty,numeric,len=11"`
	Password *string `json:"password,omitempty" validate:"omitempty,gte=6,lte=32"`
	Nickname *string `json:"nickname" validate:"omitempty,gte=2,lte=32"`
	Status   *int8   `json:"status" validate:"omitempty"`
}
