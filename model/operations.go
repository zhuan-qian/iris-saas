package model

const (
	SYSTEM_PARAMETER_FOR_OPERATIONS_CRUMBS        = "operations_crumbs"        //发现页导航
	SYSTEM_PARAMETER_FOR_OPERATIONS_SLIDE         = "operations_slide"         //发现页轮播图
	SYSTEM_PARAMETER_FOR_OPERATIONS_GUIDE         = "operations_guide"         //发现页指南
	SYSTEM_PARAMETER_FOR_OPERATIONS_HOTKEYWORD    = "operations_hotkeyword"    //搜索页热门关键词
	SYSTEM_PARAMETER_FOR_EASY_ANIMAL_POPULARKINDS = "easy_animal_popularkinds" //受欢迎的宠物种类
)

type Operations struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT"`
	Name        string `json:"name" xorm:"not null comment('配置名称') unique VARCHAR(64)"`
	Params      string `json:"params" xorm:"not null comment('参数json') TEXT"`
	Description string `json:"description" xorm:"not null comment('描述') VARCHAR(255)"`
}
