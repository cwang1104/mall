package user

import (
	"context"
	"mall/apps/user/user/userclient"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.GetUserListReq) (resp *types.GetUserListResponse, err error) {
	data := userclient.GetUserRequest{
		CurrentPage: req.CurrentPage,
		PageSize:    req.PageSize,
	}
	var response types.GetUserListResponse
	res, err := l.svcCtx.UserRpc.GetUserList(l.ctx, &data)
	if err != nil {
		logx.Errorf("getUserList userRpc failed", err)
		return nil, err
	}
	_ = copier.Copy(&response, res)
	return &response, nil
}
