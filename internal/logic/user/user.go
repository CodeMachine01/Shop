package user

import (
	"Shop/internal/consts"
	"Shop/internal/dao"
	"Shop/internal/model"
	"Shop/internal/model/do"
	"Shop/internal/service"
	"Shop/utility"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// 用户注册
func (s *sUser) Register(ctx context.Context, in model.RegisterInput) (out model.RegisterOutput, err error) {
	//处理加密盐和密码的逻辑
	UserSalt := grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt
	//插入数据返回id
	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.RegisterOutput{Id: uint(lastInsertID)}, err
}

// 修改密码
func (s *sUser) UpdatePassword(ctx context.Context, in model.UpdatePasswordInput) (out model.UpdatePasswordOutput, err error) {
	//验证密保问题
	userInfo := do.UserInfo{}
	userId := gconv.Uint(ctx.Value(consts.CtxUserId))
	err = dao.UserInfo.Ctx(ctx).WherePri(userId).Scan(&userInfo)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	//if userInfo.SecretAnswer != in.SecretAnswer {			//值一样，但类型不一样还是不一样
	if gconv.String(userInfo.SecretAnswer) != in.SecretAnswer {
		g.Dump("userInfo.SecretAnswer", userInfo.SecretAnswer)
		g.Dump("in.SecretAnswer", in.SecretAnswer)
		fmt.Printf("userInfo.SecretAnswer类型为%T，in.SecretAnswer类型为%T", userInfo.SecretAnswer, in.SecretAnswer)
		return out, errors.New(consts.ErrSecretAnswerMSg)
	}
	UserSalt := grand.S(10)
	in.UserSalt = UserSalt
	in.Password = utility.EncryptPassword(in.Password, UserSalt)
	_, err = dao.UserInfo.Ctx(ctx).WherePri(userId).Update(in)
	if err != nil {
		return model.UpdatePasswordOutput{}, err
	}
	return model.UpdatePasswordOutput{Id: userId}, nil
}
