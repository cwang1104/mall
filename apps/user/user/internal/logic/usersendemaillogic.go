package logic

import (
	"context"
	"mall/pkg/utils"
	"time"

	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"

	cache "github.com/patrickmn/go-cache"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserSendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var emailCache = cache.New(60*time.Second, 20*time.Second)

func NewUserSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSendEmailLogic {
	return &UserSendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserSendEmailLogic) UserSendEmail(in *user.UserMailRequest) (*user.UserResponse, error) {

	randNum := utils.GetRandNum(6)
	err := utils.SendEmail(in.Email, randNum)
	if err != nil {
		return &user.UserResponse{
			Code: 500,
			Msg:  "发送邮件失败",
		}, nil
	}

	emailCache.Set(in.Email, randNum, cache.DefaultExpiration)

	return &user.UserResponse{
		Code: 200,
		Msg:  "邮件发送成功",
	}, nil
}
