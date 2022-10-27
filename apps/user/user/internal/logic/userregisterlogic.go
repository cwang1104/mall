package logic

import (
	"context"
	"log"
	"mall/apps/user/model"
	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"
	"mall/pkg/utils"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *user.UserRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	emailCode, isOk := emailCache.Get(in.Email)
	if !isOk {
		log.Println("------------------------1")
		return &user.UserResponse{
			Code: 500,
			Msg:  "注册失败，请重新尝试",
		}, nil
	} else {
		if emailCode.(string) != in.Code {
			log.Println("------------------------2")
			return &user.UserResponse{
				Code: 500,
				Msg:  "验证码不正确",
			}, nil
		} else {
			userInfo, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
			if err != nil && err != model.ErrNotFound {
				log.Println("------------------------3")
				log.Println("数据库查询失败", err)
				return nil, err
			}
			if userInfo != nil {
				log.Println("------------------------4")
				return &user.UserResponse{
					Code: 500,
					Msg:  "用户已存在",
				}, nil
			}

			hashPassword, err := utils.GetHashPassword(in.Password)
			if err != nil {
				log.Println("------------------------5")
				return &user.UserResponse{
					Code: 500,
					Msg:  "加密错误，请重试",
				}, nil
			}

			addUser := model.User{
				Email:      in.Email,
				Password:   hashPassword,
				Status:     1,
				Desc:       "测试用户",
				CreateTime: time.Now(),
			}

			_, err = l.svcCtx.UserModel.Insert(l.ctx, &addUser)

			if err != nil {
				log.Println("------------------------6")
				return &user.UserResponse{
					Code: 500,
					Msg:  "注册失败",
				}, nil
			}

			return &user.UserResponse{
				Code: 200,
				Msg:  "注册成功",
			}, nil
		}
	}
}
