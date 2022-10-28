package logic

import (
	"context"
	"mall/apps/user/model"

	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *user.AdminRequest) (*user.AdminResponse, error) {
	// todo: add your logic here and delete this line
	admin, err := l.svcCtx.AdminModel.FindOneByUserName(l.ctx, in.UserName)
	if err != nil {
		if err == model.ErrNotFound {
			return &user.AdminResponse{
				Code: 500,
				Msg:  "无此用户",
			}, nil
		}
		return &user.AdminResponse{
			Code: 500,
			Msg:  "查询错误",
		}, nil
	}

	if admin.Password != in.Password {
		return &user.AdminResponse{
			Code: 500,
			Msg:  "密码错误",
		}, nil
	}

	return &user.AdminResponse{
		Code: 200,
		Msg:  "登录成功",
	}, nil
}
