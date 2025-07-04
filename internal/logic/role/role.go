package role

import (
	"Shop/internal/dao"
	"Shop/internal/model"
	"Shop/internal/service"
	"context"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 增加
func (s *sRole) Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error) {
	//插入数据返回id
	lastInsertID, err := dao.RoleInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RoleCreateOutput{RoleId: int(lastInsertID)}, err
}
