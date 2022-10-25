package user

import (
	"context"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterHandleLogic {
	return &RegisterHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterHandleLogic) RegisterHandle(req *types.RegisterRequest) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
