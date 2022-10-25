package user

import (
	"context"
	"github.com/pkg/errors"
	"mall/apps/user/user/user"
	"mall/pkg/jwtx"
	"time"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginHandleLogic {
	return &LoginHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginHandleLogic) LoginHandle(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	var loginReq user.UserRequest
	loginReq.Email = req.Email
	loginReq.Password = req.Password

	res, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &loginReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req:%+v", req)
	}

	//生成jwt token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.Email)
	if err != nil {
		return nil, err
	}
	resp.Msg = res.Msg
	resp.Code = res.Code
	resp.UserName = res.Email
	resp.Token = accessToken
	resp.AccessExpire = now + accessExpire
	return
}
