// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"Shop/internal/model"
	"context"
)

type (
	IAdmin interface {
		Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, in model.AdminUpdateInput) error
		GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
		GetAdminByNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{}
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
