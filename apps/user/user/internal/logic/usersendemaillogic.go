package logic

import (
	"context"

	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSendEmailLogic {
	return &UserSendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSendEmailLogic) UserSendEmail(in *user.UserMailRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserResponse{}, nil
}
