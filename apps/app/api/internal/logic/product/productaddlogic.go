package product

import (
	"context"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAddLogic) ProductAdd(req *types.ProductAddReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
