package logic

import (
	"context"
	"mall/apps/user/model"
	"mall/apps/user/user/userclient"
	"strconv"

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

	if err != nil {
		if err == model.ErrNotFound {
			return &user.GetUserResponse{
				Code: 200,
				Msg:  "无数据",
			}, nil
		}
		logx.Errorf("db 查询失败 %v", err)
		return &user.GetUserResponse{
			Code: 500,
			Msg:  "查询失败",
		}, nil
	}
	var resUsers []*userclient.UserInfo
	//_ = copier.Copy(&resUsers, users)
	for _, v := range users {
		userInfo := userclient.UserInfo{
			Email:       v.Email,
			Status:      strconv.Itoa(int(v.Status)),
			CreatedTime: v.CreateTime.String(),
			Desc:        v.Desc,
		}
		resUsers = append(resUsers, &userInfo)
	}

	return &user.GetUserResponse{
		Code:     200,
		Msg:      "查询成功",
		Users:    resUsers,
		Total:    int32(len(resUsers)),
		Current:  in.CurrentPage,
		PageSize: in.PageSize,
	}, nil
}
