package user

import (
	"context"
	"mall/apps/user/user/userclient"

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

	var users []types.User
	for _, v := range res.Users {
		user := types.User{
			Email:       v.Email,
			Desc:        v.Desc,
			Status:      v.Status,
			CreatedTime: v.CreatedTime,
		}
		users = append(users, user)
	}
	response.Code = res.Code
	response.Msg = res.Msg
	response.FrontUsers = users
	response.PageSize = res.PageSize
	response.CurrentPage = res.Current
	response.Total = res.Total
	return &response, nil
}
