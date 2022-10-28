package user

import (
	"context"
	"mall/apps/user/user/userclient"
	"mall/pkg/utils"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.SendEmailRequest) (resp *types.SendEmailResponse, err error) {

	if utils.VerifyEmail(req.Email) == false {
		resp = &types.SendEmailResponse{
			Code: 500,
			Msg:  "email格式错误",
		}
		return
	}

	var pbData userclient.UserMailRequest
	pbData.Email = req.Email

	res, err := l.svcCtx.UserRpc.UserSendEmail(l.ctx, &pbData)
	if err != nil {
		return nil, err
	}

	resp = &types.SendEmailResponse{
		Code: res.Code,
		Msg:  res.Msg,
	}
	return
}
