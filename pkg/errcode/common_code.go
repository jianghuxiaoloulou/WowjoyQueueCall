package errcode

var (
	Http_Success   = NewError(0, "成功")
	Http_Created   = NewError(201, "Created")
	Http_NoContent = NewError(204, "No Content")

	Http_RequestError = NewError(400, "失败")
	Http_Unauthorized = NewError(401, "Unauthorized")
	Http_Forbidden    = NewError(403, "Forbidden")
	Http_NotFound     = NewError(404, "Not Found")

	Http_Error     = NewError(500, "失败")
	Http_HeadError = NewError(600, "增加Head参数错误")
	Http_RespError = NewError(601, "获取请求结果错误")

	File_OpenError = NewError(1000, "文件打开错误")
	File_CopyError = NewError(1001, "文件拷贝错误")

	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
)
