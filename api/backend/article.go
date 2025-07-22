package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ArticleReq struct {
	g.Meta `path:"/article/add" tags:"文章" method:"post" summary:"添加文章"`
	ArticleCommonAddUpdate
}

type ArticleCommonAddUpdate struct {
	PicUrl  string `json:"pic_url"           description:"图片"`
	Title   string `json:"title" dc:"文章标题" v:"required#名称必填"`
	Desc    string `json:"desc" dc:"文章概要"`
	IsAdmin uint   `d:"1" desc:"1后台管理员发布 2前台用户发布"`
	Detail  string `json:"detail" dc:"文章详情"  v:"required#名称详情必填"`
	Praise  int    `json:"praise"  dc:"点赞数量"`
}

type ArticleRes struct {
	Id uint `json:"id"`
}

type ArticleDeleteReq struct {
	g.Meta `path:"/article/delete" method:"delete" tags:"文章" summary:"删除文章接口"`
	Id     int `v:"min:1#请选择需要删除的文章" dc:"文章id"`
}
type ArticleDeleteRes struct{}

type ArticleUpdateReq struct {
	g.Meta `path:"/article/update/{Id}" method:"post" tags:"文章" summary:"修改文章接口"`
	Id     int `json:"id"      v:"min:1#请选择需要修改的文章" dc:"文章Id"`
	ArticleCommonAddUpdate
}
type ArticleUpdateRes struct {
	Id int `json:"id"`
}

type ArticleGetListCommonReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章" summary:"查询文章列表"`
	CommonPaginationReq
}
type ArticleGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}
