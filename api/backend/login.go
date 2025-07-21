package backend

import (
	"Shop/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

//type LoginIndexRes struct {
//	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
//}

type LoginDoReq struct {
	//g.Meta   `path:"/login" method:"post" tags:"登录" summary:"登录" `
	Name     string `json:"name" v:"required#请输入账号"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

//
//// for jwt
//type LoginDoRes struct {
//	//Info interface{} `json:"info"`
//	//Referer string `json:"referer" dc:"引导客户端跳转地址"`
//	Token  string    `json:"token"`
//	Expire time.Time `json:"expire"`
//}

// for gtoken
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`
	RoleIds     string                  `json:"role_ids"`
	Permissions []entity.PermissionInfo `json:"permissions"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post" tags:"登录" summary:"刷新、续签"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post" tags:"登录" summary:"退出"`
}

type LogoutRes struct {
}
