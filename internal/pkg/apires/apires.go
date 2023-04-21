package apires

//Api接口统一返回定义
type ApiRes struct {
	// 请求id
	RequestId string `json:"RequestId"`
	//错误码/返回码
	Code     string      `json:"Code,omitempty"`
	MoreInfo interface{} `json:"MoreInfo,omitempty"`
	Status   string      `json:"Status"`
	Data     interface{} `json:"Data,omitempty"`
}

var (
	NormalSucess    = ApiRes{}
	SystemException = ApiRes{
		Code:     "IOTSYS.000001",
		Status:   "error",
		MoreInfo: "System Exception",
	}
)
