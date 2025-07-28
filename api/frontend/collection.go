package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddCollectionReq struct {
	g.Meta   `path:"/add/collection" method:"post" tags:"前台收藏" summary:"添加收藏"`
	UserId   uint `json:"user_id"    description:"用户id"`
	ObjectId int  `json:"object_id"  description:"对象id" v:"required#收藏的id必填"`
	Type     int  `json:"type"      description:"收藏类型：1商品 2文章" v:"in:1,2"` //数据校验范围约束
}

type AddCollectionRes struct {
	Id uint `json:"id"`
}

type DeleteCollectionReq struct {
	g.Meta   `path:"/delete/collection" method:"post" tags:"前台收藏" summary:"移除收藏"`
	Id       uint  `json:"id"`
	Type     uint8 `json:"type"`
	ObjectId uint  `json:"object_id"`
}
type DeleteCollectionRes struct {
	Id uint `json:"id"`
}
