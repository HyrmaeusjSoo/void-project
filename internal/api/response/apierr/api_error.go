package apierr

import "void-project/internal/api/response"

var (
	BadRequest          = response.NewJsonResError(400_00_000, "请求不合法") // 请求不合法
	Unauthorized        = response.NewJsonResError(401_00_000, "访问未授权") // 访问未授权
	Forbidden           = response.NewJsonResError(403_00_000, "无权限访问") // 无权限访问
	NotFound            = response.NewJsonResError(404_00_000, "资源未找到") // 资源未找到
	InternalServerError = response.NewJsonResError(500_00_000, "系统错误")  // 系统错误

	Unknown                  = response.NewJsonResError(500_00_001, "未知错误")               // 未知错误
	Failed                   = response.NewJsonResError(500_00_002, "请求失败")               // 请求失败
	InvalidParameter         = response.NewJsonResError(400_00_003, "无效的参数")              // 无效的参数
	MissingRequiredParameter = response.NewJsonResError(400_00_004, "必传参数缺失")             // 必传参数缺失
	OperationFailed          = response.NewJsonResError(200_00_005, "操作失败")               // 操作失败
	CreateFailed             = response.NewJsonResError(500_00_006, "创建失败")               // 创建失败
	UpdateFailed             = response.NewJsonResError(500_00_007, "更新失败")               // 更新失败
	SaveFailed               = response.NewJsonResError(500_00_008, "保存失败")               // 保存失败
	DeleteFailed             = response.NewJsonResError(500_00_009, "删除失败")               // 删除失败
	FetchFailed              = response.NewJsonResError(500_00_010, "查询失败")               // 查询失败
	AuthInvalidUserId        = response.NewJsonResError(401_00_011, "鉴权失败,无效的userId")     // 鉴权失败,无效的userId
	AuthInvalidToken         = response.NewJsonResError(401_00_012, "鉴权失败,无效的token")      // 鉴权失败,无效的token
	AuthExpired              = response.NewJsonResError(401_00_013, "鉴权失败,授权已过期")         // 鉴权失败,授权已过期
	AuthUserIdMismatch       = response.NewJsonResError(401_00_014, "鉴权失败,userId与其凭证不匹配") // 鉴权失败,userId与其凭证不匹配
	AuthClaimsExpired        = response.NewJsonResError(401_00_015, "鉴权失败,登录已失效")         // 鉴权失败,登录已失效

	// User
	MissingAccountPassword = response.NewJsonResError(400_01_001, "缺失账号或密码") // 缺失账号或密码
	AccountNotExist        = response.NewJsonResError(400_01_002, "账号不存在")   // 账号不存在
	InvalidPassword        = response.NewJsonResError(400_01_003, "密码错误")    // 密码错误
	MissingAcctPwd         = response.NewJsonResError(400_01_004, "缺失账号或密码") // 缺失账号或密码
	PasswordMismatch       = response.NewJsonResError(400_01_005, "两次密码不一致") // 两次密码不一致
	AccountExists          = response.NewJsonResError(400_01_006, "该账号已被注册") // 该账号已被注册
)
