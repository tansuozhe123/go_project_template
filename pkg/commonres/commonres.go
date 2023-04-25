package commonres

// Api接口统一返回定义
type CommonRes struct {
	// 请求id
	RequestId string `json:"RequestId"`
	//错误码/返回码
	ErrCode string `json:"ErrCode,omitempty"`
	//错误信息
	Message interface{} `json:"Message,omitempty"`
	//返回类型
	Type string `json:"Type"`
	//返回数据
	Data interface{} `json:"Data,omitempty"`
}

var (
	NormalSucess = CommonRes{
		Type: "success",
	}
	//系统异常
	SystemException = CommonRes{
		Type:    "error",
		ErrCode: "IOTSYS.000001",
		Message: "System Exception",
	}
	//您已欠费超期
	BillIsOverDue = CommonRes{
		Type:    "error",
		ErrCode: "IOTSYS.000002",
		Message: "Bill Is Over Due",
	}
	//非法参数请求
	IllegalRequest = CommonRes{
		Type:    "error",
		ErrCode: "IOTSYS.000003",
		Message: "Illegal Request",
	}
	//非法的返回值
	IllegalResponse = CommonRes{
		Type:    "error",
		ErrCode: "IOTSYS.000004",
		Message: "Illegal Response",
	}
	//账号不存在或未开通服务
	UidNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTSYS.000005",
		Message: "Uid Not Found",
	}
	//API请求方法不存在
	APINoFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000001",
		Message: "Api No Found",
	}
	//无法找到后端
	BackendNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000002",
		Message: "Backend Not Found",
	}
	//无法找到插件配置
	PluginNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000003",
		Message: "Plugin Not Found",
	}
	//无法找到后端配置
	BackedConfigNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000004",
		Message: "Backed Config Not Found",
	}
	//编排错误
	OrchestrationError = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000005",
		Message: "Orchestration Error",
	}
	//请求格式不合法或标头缺失
	InvalidRequestFormat = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000006",
		Message: "Invalid Request Format",
	}
	//请求body过大
	InvalidRequestBodyLarge = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000007",
		Message: "Invalid Request Body Large",
	}
	//请求URI过大
	InvalidRequestURILarge = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000008",
		Message: "Invalid Request URI Large",
	}
	//请求头过大
	InvalidRequestHeadersLarge = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000009",
		Message: "Invalid Request Headers Large",
	}
	//后端不可用
	BackendUnavailable = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000010",
		Message: "Backend Unavailable",
	}
	//后端超时
	BackendTimeout = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000011",
		Message: "Backend Timeout",
	}
	//APP认证信息错误
	IncorrectAppAuthInfo = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000012",
		Message: "Incorrect App Auth Info",
	}
	//app无权限访问api
	AppAccessApiDenied = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000013",
		Message: "App Access Api Denied",
	}
	//认证信息错误
	IncorrectAuthInfo = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000014",
		Message: "Incorrect Auth Info",
	}
	//不允许访问api
	APIAccessDenied = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000015",
		Message: "API Access Denied",
	}
	//tokem需要更新
	InvalidOldToken = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000016",
		Message: "Invalid Old Token",
	}
	//超出流控制限制
	ThrottlingCountRechedMax = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000017",
		Message: "Throttling Count Reched Max",
	}
	//project不可用
	ProjectUnavaliable = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000018",
		Message: "Project Unavaliable",
	}
	//调试认证信息不可用
	IncorrectDebugAuthInfo = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000019",
		Message: "Incorrect Debug Auth Info",
	}
	//无法识别客户端IP地址
	UnknownClientIPAddress = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000020",
		Message: "Unknown Client IP Address",
	}
	//ip地址不允许访问
	NotAuthorizedIPAddress = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000021",
		Message: "Not Authorized IP Address",
	}
	//后端IP地址不允许访问
	BackendIPAccessDenied = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000022",
		Message: "Backend IP Access Denied",
	}
	//内部错误
	InternalServerError = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000023",
		Message: "Internal Server Error",
	}
	//非法请求
	InvalidRequest = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000024",
		Message: "Invalid Request",
	}
	//域名解析失败
	InvalidDomainName = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000025",
		Message: "Invalid Domain Name",
	}
	//未加载api配置
	NotFoundAPIConfig = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000026",
		Message: "Not Found API Config",
	}
	//协议不被允许
	InvalidProtocol = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000027",
		Message: "Invalid Protocol",
	}
	//无法获取管理租户
	InvalidAdminToken = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000028",
		Message: "Invalid Admin Token",
	}
	//找不到vpc后端
	VPCNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000029",
		Message: "VPC Not Found",
	}
	//没有可链接的后端
	NoAvaliableAPI = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000030",
		Message: "No Avaliable API",
	}
	//后端端口未找到
	APIPortNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000031",
		Message: "API Port Not Found",
	}
	//api调用自身
	InvaildAPISelfRequest = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000032",
		Message: "Invaild API Self Request",
	}
	//计算后端签名失败
	SignatureCalculationFailed = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000033",
		Message: "Signature Calculation Failed",
	}
	//服务在当前region不可访问
	APINotAccessibleRegion = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000034",
		Message: "API Not Accessible Region",
	}
	//该用户在当前region中被禁用
	UserAccessForbiddenRegion = CommonRes{
		Type:    "error",
		ErrCode: "IOTGW.000035",
		Message: "User Access Forbidden Region",
	}
	//数组中包含非法元素
	InvalidElementInArray = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000001",
		Message: "Invalid Element In Array",
	}
	//参数格式错误
	InvalidFormattedParameter = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000002",
		Message: "Invalid Formatted Parameter",
	}
	//设备数组中包含有不存在的设备
	NotFoundDeviceInDeviceNameArray = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000003",
		Message: "Not Found Device In DeviceName Array",
	}
	//参数不能为空
	NullParameter = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000004",
		Message: "Null Parameter",
	}
	//账号信息不能为空
	AccountInfoCannotNull = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000005",
		Message: "Account Info Cannot Null",
	}
	//未通过实名认证，不符合购买条件
	AccountNameIsNotReal = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000006",
		Message: "Account Name Is Not Real",
	}
	//已开通云平台服务
	AlreadyOpenLinkPlatForm = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000007",
		Message: "Already Open Link PlatForm",
	}
	//当前账户不支持该api
	APINotSupportedInAccount = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000008",
		Message: "API Not Supported In Account",
	}
	//欠费用户
	ArrearageAccount = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000009",
		Message: "Arrearage Account",
	}
	//鉴权失败，原因可能是入参的设备不属于当前账号
	AuthActionPermissionDeny = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000010",
		Message: "Auth Action Permission Deny",
	}
	//无权操作此设备或产品
	NotAuthOperateProductOrDevice = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000011",
		Message: "Not Auth Operate Product Or Device",
	}
	//创建服务账号关联角色失败
	CreateServiceLinkedRoleFailed = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000012",
		Message: "Create Service Linked Role Failed",
	}
	//数组参数不能为空
	EmptyListParam = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000013",
		Message: "Empty List Param",
	}
	//Region 信息不能为空
	EmptyRegion = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000014",
		Message: "Empty Region",
	}
	//Project 信息不能为空
	EmptyProject = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000015",
		Message: "Empty Project",
	}
	//rolename不能为空
	EmptyRoleName = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000016",
		Message: "Empty RoleName",
	}
	//文件不不存在
	FileNotExist = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000017",
		Message: "File Not Exist",
	}
	//触发系统限流
	FlowControlTriggered = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000018",
		Message: "Flow Control Triggered",
	}
	//文件业务类型无效
	IllegalFileBizCode = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000019",
		Message: "Illegal File BizCode",
	}
	//文件格式无效
	IllegalFileFormat = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000020",
		Message: "Illegal File Format",
	}
	//非法产品校验数据类型
	IllegalProductValidateType = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000021",
		Message: "Illegal Product Validate Type",
	}
	//物模型数据类型错误
	IllegalTLSValidateType = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000022",
		Message: "Illegal TLS Validate Type",
	}
	//不支持的访问方式
	InvalidCallerTypeError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000023",
		Message: "Invalid Caller Type Error",
	}
	//参数格式错误
	ParameterFormattedError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000024",
		Message: "Parameter Formatted Error",
	}
	//标签key格式不合法
	InvalidFormattedTagKey = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000025",
		Message: "Invalid Formatted Tag Key",
	}
	//标签value格式不合法
	InvalidFormattedTagValue = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000026",
		Message: "Invalid Formatted Tag Value",
	}
	//分页大小或分页页号不合法
	InvalidPageParams = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000027",
		Message: "Invalid Page Params",
	}
	//不合法租户
	InvalidTenant = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000028",
		Message: "Invalid Tenant",
	}
	//不合法账号或未开通服务
	InvalidUid = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000029",
		Message: "Invalid Uid",
	}
	//禁止访问接口，失败次数过多
	MethodInvokeNotPermitted = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000030",
		Message: "Method Invoke Not Permitted",
	}
	//不允许更新默认TSL模块名称
	NotAllowUpdateTLSModelName = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000031",
		Message: "Not Allow Update TLS Model Name",
	}
	//原数据类型不允许更新
	NotAllowUpdateValidateType = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000032",
		Message: "Not Allow Update Validate Type",
	}
	//参数不能为空
	EmptyParameter = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000033",
		Message: "Empty Parameter",
	}
	//产品校验数据类型为空
	NullProductValidateType = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000034",
		Message: "Null Product Validate Type",
	}
	//开通服务失败
	OpenServiceError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000035",
		Message: "Open Service Error",
	}
	//角色授权策略错误导致没有权限访问资源
	PolicyNoPermission = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000036",
		Message: "Policy No Permission",
	}
	//属性不存在
	PropertyNotFound = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000037",
		Message: "Property Not Found",
	}
	//设备不存在或未激活
	QueryDeviceActionError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000038",
		Message: "Query Device Action Error",
	}
	//查询设备属性失败
	QueryDevicePropertyActionError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000039",
		Message: "Query Device Property Action Error",
	}
	//批量查询设备失败
	QueryManyDevicesActionError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000040",
		Message: "Query Many Devices Action Error",
	}
	//查询产品失败
	QueryProductActionError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000041",
		Message: "Query Product Action Error",
	}
	//查询产品总数失败
	QueryProductCountActionError = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000042",
		Message: "Query Product Count Action Error",
	}
	//查询产品数据校验类型失败
	QueryProductValidationFailed = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000043",
		Message: "Query Product Validation Failed",
	}
	//获取用户信息失败
	QueryUserInfoFailed = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000044",
		Message: "Query User Info Failed",
	}
	//请求禁止
	RequestForbidden = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000045",
		Message: "Request Forbidden",
	}
	//服务账号关联角色已存在
	ServiceLinkedRoleAlreadyExists = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000046",
		Message: "Service Linked Role Already Exists",
	}
	//物模型功能模块名称重复
	DuplicatedTLSModelName = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000047",
		Message: "Duplicated TLS Model Name",
	}
	//物模型功能模块不存在
	TSLModelNameNotExist = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000048",
		Message: "TSL Model Name Not Exist",
	}
	//物模型功能模块个数超限
	TSLModelCountOverLimit = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000049",
		Message: "TSL Model Count Over Limit",
	}
	//物模型功能模块数据量过大
	TSLModelObjectSizeOverLimit = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000050",
		Message: "TSL Model Object Size Over Limit",
	}
	//超出更新物模型最大数量
	UpdateTSLModelOverLimit = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000051",
		Message: "Update TSL Model Over Limit",
	}
	//请求频率过高
	TooManyRequest = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000052",
		Message: "Too Many Request",
	}
	//待更新数据类型与当前值一致
	ValidateTypeIsSame = CommonRes{
		Type:    "error",
		ErrCode: "IOTCOM.000053",
		Message: "Validate Type Is Same",
	}
)
