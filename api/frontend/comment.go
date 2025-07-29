package frontend

import "github.com/gogf/gf/v2/frame/g"

type AddCommentReq struct {
	g.Meta   `path:"/add/comment" method:"post" tags:"前台评论" summary:"添加评论"`
	Type     int    `json:"type"      description:"评论类型 1商品 2文章"`
	ObjectId int    `json:"object_id"  description:"评论对象id 方便后期扩展"`
	ParentId uint   `json:"parent_id" dc:"父级评论id"`
	Content  string `json:"content" v:"required#评论必填"`
}

type AddCommentRes struct {
	Id uint `json:"id"`
}

// 数据库唯一索引限制不允许对同一个内容发相同的评论
type DeleteCommentReq struct {
	g.Meta `path:"/delete/comment" method:"post" tags:"前台评论" summary:"移除评论"`
	Id     uint `json:"id"`
}
type DeleteCommentRes struct {
	Id uint `json:"id"`
}

type ListCommentReq struct {
	g.Meta `path:"/comment/list" method:"post" tags:"前台评论" summary:"评论列表"`
	Type   uint8 `json:"type" v:"in:0,1,2" dc:"评论类型"`
	CommonPaginationReq
}

type ListCommentRes struct {
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
	List  interface{} `json:"list" description:"列表"`
}

type ListCommentItem struct {
	Id       int         `json:"id"        description:""`
	UserId   int         `json:"user_id"    description:"用户id"`
	ObjectId int         `json:"object_id"  description:"对象id"`
	Type     int         `json:"type"      description:"评论类型：1商品 2文章"`
	Goods    interface{} `json:"goods"`
	Article  interface{} `json:"article"`
}
