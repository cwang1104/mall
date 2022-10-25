package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/apps/app/api/internal/config"
	"mall/apps/user/user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
