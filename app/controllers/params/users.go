package params

type UsersGet struct {
	Keyword string `json:"keyword"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
}

type UsersPost struct {
	Account string `json:"account" validate:"required,len=11"`
	Code    string `json:"code" validate:"required,len=4"`
}

type UsersModify struct {
	Password        *string `json:"password" validate:"omitempty,max=24,min=6"`
	ConfirmPassword *string `json:"confirmPassword" validate:"omitempty,max=24,min=6"`
	Nickname        *string `json:"nickname" validate:"omitempty,max=32,min=2"`
	Avatar          *string `json:"avatar" validate:"omitempty,url"`
	BornAt          *string `json:"bornAt" validate:"omitempty,len=10"`
	GotPetAt        *string `json:"gotPetAt" validate:"omitempty,len=10"`
	PushToken       *string `json:"pushToken" validate:"omitempty,max=128,min=32"`
	DynamicLock     *byte   `json:"dynamicLock" validate:"omitempty,max=1,min=0"`
	LockTimeTill    *string `json:"lockTimeTill" validate:"omitempty,datetime=2016-01-02 15:04:05"`
	LastLoginAt     *string `json:"lastLoginAt" validate:"omitempty,len=10"`
	NewbieGuided    *byte   `json:"newbieGuided" validate:"omitempty,max=1,min=0"`
	Status          *int8   `json:"status" validate:"omitempty,numeric"`
	TalkStatus      *int8   `json:"talkStatus" validate:"omitempty,numeric,oneof=0 1"`
	Sex             *byte   `json:"sex" validate:"numeric,oneof=0 1 2"`
	OrgId           *int    `json:"orgId,omitempty" validate:"omitempty,numeric"`
	Saleable        *int8   `json:"saleable" validate:"omitempty,numeric,oneof=0 1"`
}
