package common

//常规定义
const (
	TOKEN_NAME_IN_CLIENT       = "JWT"
	TOKEN_NAME_IN_BACKEND      = "Authorization"
	RBAC_URL_VAL_SYMBOL        = "$v"
	PARAMS_VALUE_USER          = "user"
	GOODS_FOLLOW_STATUS_TRUE   = 1 //关注状态为关注
	GOODS_FOLLOW_STATUS_FALSE  = 0 // 关注状态为不关注
	GOODS_COLLECT_STATUS_TRUE  = 1 //收藏状态为收藏
	GOODS_COLLECT_STATUS_FALSE = 0 //收藏状态为不收藏
)

//资源目录定义
const (
	PUBLIC_RESOURCE = "public/resource"
)

//返回码定义
const (
	RESPONSE_EXPLAIN_BY_MSG = "0001"

	// ----- 底层错误区块 Begin -----
	RESPONSE_CONTACTS_BROKEN = "1001" //长连接错误码
	RESPONSE_ONLINE_FAILD    = "1002"
	RESPONSE_OFFLINE_FAILD   = "1003"
	RESPONSE_SERVER_ERROR    = "1004" //常规服务错误
	// ----- 底层错误区块 End -----

	// ----- 校验错误区块 Begin -----
	RESPONSE_INVALID_TOKEN = "2000"
	RESPONSE_REQUEST_ERROR = "2001" //请求参数错误

	RESPONSE_CRETE_ORDER_ERR = "2010"
	// ----- 校验错误区块 End -----

	// ----- 数据逻辑错误区块 Begin -----
	RESPONSE_SQL_ERROR        = "3000"
	RESPONSE_SQL_INSERT_FAILD = "3001"
	RESPONSE_SQL_UPDATE_FAILD = "3002"
	// ----- 数据逻辑错误区块 End -----

)
