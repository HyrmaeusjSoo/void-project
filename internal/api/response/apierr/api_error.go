package apierr

import "void-project/internal/api/response"

var (
	BadRequest          = response.NewJsonResponseError(400_00_000, "请求不合法") // 请求不合法
	Unauthorized        = response.NewJsonResponseError(401_00_000, "访问未授权") // 访问未授权
	Forbidden           = response.NewJsonResponseError(403_00_000, "无权限访问") // 无权限访问
	NotFound            = response.NewJsonResponseError(404_00_000, "资源未找到") // 资源未找到
	InternalServerError = response.NewJsonResponseError(500_00_000, "系统错误")  // 系统错误

	Unknown            = response.NewJsonResponseError(500_00_001, "未知错误")               // 未知错误
	Failed             = response.NewJsonResponseError(500_00_002, "请求失败")               // 请求失败
	InvalidParameter   = response.NewJsonResponseError(400_00_003, "无效的参数")              // 无效的参数
	OperationFailed    = response.NewJsonResponseError(200_00_004, "操作失败")               // 操作失败
	CreateFailed       = response.NewJsonResponseError(500_00_005, "创建失败")               // 创建失败
	UpdateFailed       = response.NewJsonResponseError(500_00_006, "更新失败")               // 更新失败
	SaveFailed         = response.NewJsonResponseError(500_00_007, "保存失败")               // 保存失败
	DeleteFailed       = response.NewJsonResponseError(500_00_008, "删除失败")               // 删除失败
	FetchFailed        = response.NewJsonResponseError(500_00_009, "查询失败")               // 查询失败
	AuthInvalidUserId  = response.NewJsonResponseError(401_00_010, "鉴权失败,无效的userId")     // 鉴权失败,无效的userId
	AuthInvalidToken   = response.NewJsonResponseError(401_00_011, "鉴权失败,无效的token")      // 鉴权失败,无效的token
	AuthExpired        = response.NewJsonResponseError(401_00_012, "鉴权失败,授权已过期")         // 鉴权失败,授权已过期
	AuthUserIdMismatch = response.NewJsonResponseError(401_00_013, "鉴权失败,userId与其凭证不匹配") // 鉴权失败,userId与其凭证不匹配

	// User
	MissingAccountPassword = response.NewJsonResponseError(400_01_001, "缺失账号或密码") // 缺失账号或密码
	AccountNotExist        = response.NewJsonResponseError(400_01_002, "账号不存在")   // 账号不存在
	InvalidPassword        = response.NewJsonResponseError(400_01_003, "密码错误")    // 密码错误
	MissingAcctPwd         = response.NewJsonResponseError(400_01_004, "缺失账号或密码") // 缺失账号或密码
	PasswordMismatch       = response.NewJsonResponseError(400_01_005, "两次密码不一致") // 两次密码不一致
	AccountExists          = response.NewJsonResponseError(400_01_006, "该账号已被注册") //该账号已被注册
)
