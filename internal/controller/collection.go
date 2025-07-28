package controller

import (
	"Shop/api/frontend"
	"Shop/internal/model"
	"Shop/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

// Collection 内容管理
var Collection = cCollection{}

type cCollection struct{}

func (a *cCollection) Add(ctx context.Context, req *frontend.AddCollectionReq) (res *frontend.AddCollectionRes, err error) {
	data := model.AddCollectionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.Collection().AddCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.AddCollectionRes{Id: out.Id}, nil
}

func (a *cCollection) Delete(ctx context.Context, req *frontend.DeleteCollectionReq) (res *frontend.DeleteCollectionRes, err error) {
	data := model.DeleteCollectionInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	collection, err := service.Collection().DeleteCollection(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.DeleteCollectionRes{Id: collection.Id}, nil
}
