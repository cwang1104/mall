package seckill

import (
	"context"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeckillResLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSeckillResLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillResLogic {
	return &GetSeckillResLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSeckillResLogic) GetSeckillRes() (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
