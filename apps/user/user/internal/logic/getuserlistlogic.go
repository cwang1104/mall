package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"mall/apps/user/model"
	"mall/apps/user/user/userclient"

	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserListLogic) GetUserList(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	offset := (in.CurrentPage - 1) * in.PageSize
	users, err := l.svcCtx.UserModel.FindAll(l.ctx, in.PageSize, offset)
	fmt.Printf("users\n%+v", users)
	if err != nil {
		if err != model.ErrNotFound {
			return &user.GetUserResponse{
				Code: 200,
				Msg:  "无数据",
			}, nil
		}
		logx.Errorf("db 查询失败", err)
		return &user.GetUserResponse{
			Code: 500,
			Msg:  "查询失败",
		}, nil
	}
	var resUsers []*userclient.UserInfo
	_ = copier.Copy(&resUsers, users)

	return &user.GetUserResponse{
		Code:     200,
		Msg:      "查询成功",
		Users:    resUsers,
		Total:    10,
		Current:  in.CurrentPage,
		PageSize: in.PageSize,
	}, nil
}
