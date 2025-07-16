package consts

const (
	ProjectName              = "开源电商项目"
	ProjectUsage             = "学习使用"
	ProjectBrief             = "start http server"
	Version                  = "v0.2.0"             // 当前服务版本(用于模板展示)
	CaptchaDefaultName       = "CaptchaDefaultName" // 验证码默认存储空间名称
	ContextKey               = "ContextKey"         // 上下文变量存储键名，前后端系统共享
	FileMaxUploadCountMinute = 10                   // 同一用户1分钟之内最大上传数量
	GtokenAdminPrefix        = "Admin:"             //Gtoken管理后台前缀区分
	CtxAdminId               = "CtxAdminId"
	CtxAdminName             = "CtxAdminName"
	CtxAdminIsAdmin          = "CtxAdminIsAdmin"
	CtxAdminRoleIds          = "CtxAdminRoleIds"
	CacheModeRedis           = 2
	BackendServerName        = "开源电商系统"
	MultiLogin               = true
	GtokenExpireIn           = 10 * 24 * 60 * 60
	//统一错误管理提示
	CodeMissingParameterMSg = "请检查是否缺少参数"
	ErrLoginFaulMSg         = "登录失败，账号或密码错误"
)
