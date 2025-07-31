package model

import (
	"Shop/internal/model/do"
	"Shop/internal/model/entity"
)

// GoodsCreateUpdateBase 创建/修改内容基类
type GoodsCreateUpdateBase struct {
	PicUrl           string
	Name             string
	Price            int
	Level1CategoryId int
	Level2CategoryId int
	Level3CategoryId int
	Brand            string
	Stock            int
	Sale             int
	Tags             string
	DetailInfo       string
}

// GoodsCreateInput 创建内容
type GoodsCreateInput struct {
	GoodsCreateUpdateBase
}

// GoodsCreateOutput 创建内容返回结果
type GoodsCreateOutput struct {
	Id uint `json:"id"`
}

// GoodsUpdateInput 修改内容
type GoodsUpdateInput struct {
	GoodsCreateUpdateBase
	Id int
}

// GoodsGetListInput 获取内容列表
type GoodsGetListInput struct {
	Page int // 分页号码
	Size int // 分页数量，最大50
	Sort int // 排序类型
}

// GoodsGetListOutput 查询列表结果
type GoodsGetListOutput struct {
	List  []GoodsGetListOutputItem `json:"list" description:"列表"`
	Page  int                      `json:"page" description:"分页码"`
	Size  int                      `json:"size" description:"分页数量"`
	Total int                      `json:"total" description:"数据总数"`
}

type GoodsGetListOutputItem struct {
	entity.GoodsInfo
	//Id        int         `json:"id"` // 自增ID
	//UserId    uint        `json:"user_id"`
	//CouponId  uint        `json:"coupon_id"`
	//Status    uint8       `json:"status"`
	//CreatedAt *gtime.Time `json:"created_at"` // 创建时间
	//UpdatedAt *gtime.Time `json:"updated_at"` // 修改时间
}

type GoodsDetailInput struct {
	Id uint
}

// todo 替换从do层的模型结构体
type GoodsDetailOutput struct {
	do.GoodsInfo
	Options   []do.GoodsOptionsInfo `orm:"with:goods_id=id"` //规格 sku
	Comments  []do.CommentInfo      `orm:"with:object_id=id, where:type=1"`
	IsCollect bool
}
