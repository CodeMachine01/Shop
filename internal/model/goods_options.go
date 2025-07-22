package model

import "Shop/internal/model/entity"

// GoodsOptionsCreateUpdateBase 创建/修改内容基类
type GoodsOptionsCreateUpdateBase struct {
	GoodsId uint
	PicUrl  string
	Name    string
	Price   uint
	Stock   uint
}

// GoodsOptionsCreateInput 创建内容
type GoodsOptionsCreateInput struct {
	GoodsOptionsCreateUpdateBase
}

// GoodsOptionsCreateOutput 创建内容返回结果
type GoodsOptionsCreateOutput struct {
	Id uint `json:"id"`
}

// GoodsOptionsUpdateInput 修改内容
type GoodsOptionsUpdateInput struct {
	GoodsOptionsCreateUpdateBase
	Id int
}

// GoodsOptionsGetListInput 获取内容列表
type GoodsOptionsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型
}

// GoodsOptionsGetListOutput 查询列表结果
type GoodsOptionsGetListOutput struct {
	List  []GoodsOptionsGetListOutputItem `json:"list" description:"列表"`
	Page  int                             `json:"page" description:"分页码"`
	Size  int                             `json:"size" description:"分页数量"`
	Total int                             `json:"total" description:"数据总数"`
}

type GoodsOptionsGetListOutputItem struct {
	entity.GoodsOptionsInfo
}
