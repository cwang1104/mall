package user

import (
	"context"
	"log"
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
	var response types.CommonResp
	if !isOK {
		response.Code = 500
		response.Msg = "邮箱格式不正确"
		logx.Errorf("邮箱格式不正确,err:%v", req.Email)
		return &response, nil
	}
	if req.Password != req.Repassword {
		response.Code = 500
		response.Msg = "两次输入密码不一致"
		log.Println("密码输入错误")
		return &response, nil
	}
	pbData := user.UserRequest{
		Email:      req.Email,
		Code:       req.Captche,
		Password:   req.Password,
		Repassword: req.Repassword,
	}
	res, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &pbData)
	if err != nil {
		response.Code = 500
		response.Msg = "注册失败，请重试"
		log.Println("注册失败，err", err)
		return &response, err
	}
	response.Code = res.Code
	response.Msg = res.Msg

	return &response, nil
}
