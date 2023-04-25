package apires

import "go_project/pkg/commonres"

//Api接口统一返回定义

var (
	//产品不存在
	NoProduct = commonres.CommonRes{
		ErrCode: "ServiceProduct.000001",
		Type:    "error",
		Message: "No Product",
	}
)
