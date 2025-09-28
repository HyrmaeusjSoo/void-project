package apierr

import "void-project/internal/api/response"

// 公共Response错误信息， 共通的都放在这里集中管理

var (
	// 原始http status错误

	BadRequest          = response.NewError(400_00_000, "请求不合法") // 请求不合法
	UnAuthorized        = response.NewError(401_00_000, "访问未授权") // 访问未授权
	UnLogin             = response.NewError(401_00_000, "未登录")   // 访问未登录
	Forbidden           = response.NewError(403_00_000, "无权限访问") // 无权限访问
	NotFound            = response.NewError(404_00_000, "资源未找到") // 资源未找到
	InternalServerError = response.NewError(500_00_000, "系统错误")  // 系统错误

	// 通用错误

	Unknown                  = response.NewError(500_00_001, "未知错误")                // 未知错误
	Failed                   = response.NewError(500_00_002, "请求失败")                // 请求失败
	AuthInvalidUserId        = response.NewError(401_00_003, "鉴权失败,无效的user_id")     // 鉴权失败,无效的user_id
	AuthInvalidToken         = response.NewError(401_00_004, "鉴权失败,无效的token")       // 鉴权失败,无效的token
	AuthExpired              = response.NewError(401_00_005, "鉴权失败,授权已过期")          // 鉴权失败,授权已过期
	AuthUserIdMismatch       = response.NewError(401_00_006, "鉴权失败,user_id与其凭证不匹配") // 鉴权失败,user_id与其凭证不匹配
	AuthClaimsExpired        = response.NewError(401_00_007, "鉴权失败,登录已失效")          // 鉴权失败,登录已失效
	InvalidParameter         = response.NewError(400_00_008, "无效的参数")               // 无效的参数
	MissingRequiredParameter = response.NewError(400_00_009, "必传参数缺失")              // 必传参数缺失
	OperationFailed          = response.NewError(400_00_010, "操作失败")                // 操作失败
	CreateFailed             = response.NewError(400_00_011, "创建失败")                // 创建失败
	UpdateFailed             = response.NewError(400_00_012, "更新失败")                // 更新失败
	SaveFailed               = response.NewError(400_00_013, "保存失败")                // 保存失败
	DeleteFailed             = response.NewError(400_00_014, "删除失败")                // 删除失败
	FetchFailed              = response.NewError(400_00_015, "查询失败")                // 查询失败
	InvalidPaginateParameter = response.NewError(400_00_016, "无效的分页参数")             // 无效的分页参数
	FileUploadFailed         = response.NewError(400_00_017, "文件上传失败")              // 文件上传失败
	RecordNotFound           = response.NewError(400_00_018, "记录未找到")               // 记录未找到
	DirNotExist              = response.NewError(400_00_019, "目录不存在")               // 目录不存在
	FileNotExist             = response.NewError(400_00_020, "文件不存在")               // 文件不存在
	InvalidPath              = response.NewError(400_00_021, "无效的路径")               // 无效的路径

	// User module

	MissingAcctPwd   = response.NewError(400_01_001, "缺失账号或密码") // 缺失账号或密码
	AccountNotExist  = response.NewError(400_01_002, "账号不存在")   // 账号不存在
	InvalidPassword  = response.NewError(400_01_003, "密码错误")    // 密码错误
	MissingPwd       = response.NewError(400_01_004, "缺失密码")    // 缺失密码
	PasswordMismatch = response.NewError(400_01_005, "两次密码不一致") // 两次密码不一致
	AccountExists    = response.NewError(400_01_006, "该账号已被注册") // 该账号已被注册
	UserNotExist     = response.NewError(400_01_007, "用户不存在")   // 用户不存在
	RegisterFailed   = response.NewError(400_01_008, "注册失败")    // 注册失败
	LoginFailed      = response.NewError(400_01_009, "登录失败")    // 登录失败
)
