package user

import (
	"context"
	"mall/apps/user/user/user"
	"mall/pkg/utils"

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
	isOK := utils.VerifyEmail(req.Email)

	if !isOK {
		resp.Code = 500
		resp.Msg = "邮箱格式不正确"
		return
	}

	if req.Password != req.Repassword {
		resp.Code = 500
		resp.Msg = "两次输入密码不一致"
		return
	}
	pbData := user.UserRequest{
		Email:      req.Email,
		Code:       req.Catpche,
		Password:   req.Password,
		Repassword: req.Repassword,
	}
	res, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &pbData)
	if err != nil {
		resp.Code = 500
		resp.Msg = "注册失败，请重试"
		return resp, err
	}
	resp.Code = res.Code
	resp.Msg = res.Msg

	return
}
