package user

import (
	"context"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginHandleLogic {
	return &LoginHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginHandleLogic) LoginHandle(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
