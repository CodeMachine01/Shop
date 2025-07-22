package controller

import (
	"Shop/api/frontend"
	"Shop/internal/model"
	"Shop/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var User = cUser{}

type cUser struct{}

func (a *cUser) Register(ctx context.Context, req *frontend.RegisterReq) (res *frontend.RegisterRes, err error) {
	data := model.RegisterInput{}
	err = gconv.Struct(req, &data)
	if err != nil {
		return nil, err
	}
	out, err := service.User().Register(ctx, data)
	if err != nil {
		return nil, err
	}
	return &frontend.RegisterRes{Id: out.Id}, nil
}
