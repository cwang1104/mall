package user

import (
	"context"
	"mall/apps/user/user/userclient"
	"mall/pkg/jwtx"
	"time"

	"mall/apps/app/api/internal/svc"
	"mall/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.AdminLoginReq) (resp *types.AdminLoginResponse, err error) {
	data := userclient.AdminRequest{
		UserName: req.UserName,
		Password: req.Password,
	}
	res, err := l.svcCtx.UserRpc.AdminLogin(l.ctx, &data)
	if err != nil {
		logx.Errorf("userRpc error", err)
		return nil, err
	}
	var response types.AdminLoginResponse
	response.Code = res.Code
	response.Msg = res.Msg
	if res.Code == 200 {
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
		accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.UserName)
		if err != nil {
			logx.Errorf("get jwt token failed, error:%s", err.Error())
			return nil, err
		}
		response.AdminToken = accessToken
		response.AccessExpire = accessExpire
		response.UserName = res.UserName
	}
	return &response, nil
}
