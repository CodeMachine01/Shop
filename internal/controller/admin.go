package controller

import (
	"Shop/api/backend"
	"context"

	"Shop/internal/model"
	"Shop/internal/service"
)

// Admin 内容管理
var Admin = cAdmin{}

type cAdmin struct{}

func (a *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{AdminId: out.AdminId}, nil
}

func (a *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

func (a *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return &backend.AdminUpdateRes{Id: req.Id}, nil
}

// Index rotation list
func (a *cAdmin) Index(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminGetListCommonRes{List: getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total}, nil
}
