package svc

import (
	"mall/apps/user/user/internal/config"
	"mall/apps/user/user/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
