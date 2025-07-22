package controller

import (
	"Shop/api/backend"
	"Shop/internal/consts"
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"Shop/internal/model"
	"Shop/internal/service"
)

// Article 内容管理
var Article = cArticle{}

type cArticle struct{}

func (a *cArticle) Create(ctx context.Context, req *backend.ArticleReq) (res *backend.ArticleRes, err error) {
	data := model.ArticleCreateInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	//获取当前登录用户
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	out, err := service.Article().Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleRes{Id: out.Id}, nil
}

func (a *cArticle) Delete(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().Delete(ctx, req.Id)
	return
}

func (a *cArticle) Update(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	data := model.ArticleUpdateInput{}
	err = gconv.Struct(req, &data)
	data.UserId = gconv.Int(ctx.Value(consts.CtxAdminId))
	if err != nil {
		return nil, err
	}
	err = service.Article().Update(ctx, data)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleUpdateRes{Id: req.Id}, nil
}

// Index position list
func (a *cArticle) List(ctx context.Context, req *backend.ArticleGetListCommonReq) (res *backend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetList(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
