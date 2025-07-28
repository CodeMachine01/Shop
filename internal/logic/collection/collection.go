package collection

import (
	"Shop/internal/consts"
	"Shop/internal/dao"
	"Shop/internal/model"
	"Shop/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCollection struct{}

func init() {
	service.RegisterCollection(New())
}

func New() *sCollection {
	return &sCollection{}
}

func (*sCollection) AddCollection(ctx context.Context, in model.AddCollectionInput) (res *model.AddCollectionOutput, err error) {
	in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
	id, err := dao.CollectionInfo.Ctx(ctx).InsertAndGetId(in)
	if err != nil {
		return &model.AddCollectionOutput{}, err
	}
	return &model.AddCollectionOutput{
		Id: gconv.Uint(id),
	}, nil
}

// 兼容处理：优先根据收藏id删除，收藏id为0 再根据对象id和type删除
func (*sCollection) DeleteCollection(ctx context.Context, in model.DeleteCollectionInput) (res *model.DeleteCollectionOutput, err error) {
	//优先根据收藏id删除
	if in.Id != 0 {
		_, err := dao.CollectionInfo.Ctx(ctx).WherePri(in.Id).Delete()
		if err != nil {
			return nil, err
		}
		return &model.DeleteCollectionOutput{Id: gconv.Uint(in.Id)}, nil
	} else {
		//收藏id为0：再根据对象id和类型type删除
		in.UserId = gconv.Uint(ctx.Value(consts.CtxUserId))
		id, err := dao.CollectionInfo.Ctx(ctx).OmitEmpty().
			Where(in).Delete()
		if err != nil {
			return &model.DeleteCollectionOutput{}, err
		}
		return &model.DeleteCollectionOutput{
			Id: gconv.Uint(id),
		}, nil
	}
}

// 列表
// GetList 查询内容列表
func (*sCollection) GetList(ctx context.Context, in model.CollectionListInput) (out *model.CollectionListOutput, err error) {
	var (
		m = dao.CollectionInfo.Ctx(ctx)
	)
	out = &model.CollectionListOutput{
		Page: in.Page,
		Size: in.Size,
		List: []model.CollectionListOutputItem{}, //提前实例化避免查不到数据返回为null而不是空数组
	}
	// 翻页查询
	listModel := m.Page(in.Page, in.Size)
	//条件查询
	if in.Type != 0 {
		listModel = listModel.Where(dao.CollectionInfo.Columns().Type, in.Type)
	}
	//优化：优先查询count，而不是查sql结果赋值到结构体中
	out.Total, err = listModel.Count()
	if err != nil {
		return out, err
	}
	if out.Total == 0 {
		return out, nil
	}
	//进一步优化：只查询相关的模型关联
	if in.Type == consts.CollectionTypeGoods {
		if err := listModel.With(model.GoodsItem{}).Scan(&out.List); err != nil {
			return out, nil
		}
	} else if in.Type == consts.CollectionTypeArticle {
		if err := listModel.With(model.ArticleItem{}).Scan(&out.List); err != nil {
			return out, nil
		}
	} else {
		if err := listModel.WithAll().Scan(&out.List); err != nil {
			return out, err
		}
	}
	return
}
