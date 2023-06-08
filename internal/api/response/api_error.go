package response

var ApiError = map[string]JsonResponse{
	"ApiErrBadRequest":          NewJsonResponseError(40000000, "请求不合法"),
	"ApiErrUnauthorized":        NewJsonResponseError(40100000, "认证失败"),
	"ApiErrForbidden":           NewJsonResponseError(40300000, "授权失败"),
	"ApiErrNotFound":            NewJsonResponseError(40400000, "资源未找到"),
	"ApiErrInternalServerError": NewJsonResponseError(50000000, "系统错误"),
}

var (
	ApiErrBadRequest          = NewJsonResponseError(40000000, "请求不合法")
	ApiErrUnauthorized        = NewJsonResponseError(40100000, "认证失败")
	ApiErrForbidden           = NewJsonResponseError(40300000, "授权失败")
	ApiErrNotFound            = NewJsonResponseError(40400000, "资源未找到")
	ApiErrInternalServerError = NewJsonResponseError(50000000, "系统错误")
)
