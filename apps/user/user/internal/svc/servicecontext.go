package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/apps/user/model"
	"mall/apps/user/user/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	AdminModel model.AdminModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserModel(sqlx.NewMysql(c.Mysql.DataSource), c.RedisCache),
		AdminModel: model.NewAdminModel(sqlx.NewMysql(c.Mysql.DataSource), c.RedisCache),
	}
}
