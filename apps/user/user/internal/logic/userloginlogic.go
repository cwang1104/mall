package logic

import (
	"context"
	"mall/apps/user/model"
	"mall/pkg/utils"

	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.UserRequest) (*user.UserResponse, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if err != nil && err == model.ErrNotFound {
		return &user.UserResponse{
			Code: 500,
			Msg:  "用户不存在",
		}, nil
	} else if err != nil && err != model.ErrNotFound {
		return &user.UserResponse{
			Code: 500,
			Msg:  "查询错误",
		}, nil
	}

	err = utils.CheckPassword(in.Password, userInfo.Password)
	if err != nil || in.Password != in.Repassword {
		return &user.UserResponse{
			Code: 500,
			Msg:  "密码错误",
		}, nil
	}

	return &user.UserResponse{
		Code: 200,
		Msg:  "登录成功",
	}, nil
}
