// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"mall/apps/user/user/internal/logic"
	"mall/apps/user/user/internal/svc"
	"mall/apps/user/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserRegister(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	l := logic.NewUserRegisterLogic(ctx, s.svcCtx)
	return l.UserRegister(in)
}

func (s *UserServer) UserSendEmail(ctx context.Context, in *user.UserMailRequest) (*user.UserResponse, error) {
	l := logic.NewUserSendEmailLogic(ctx, s.svcCtx)
	return l.UserSendEmail(in)
}

func (s *UserServer) UserLogin(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}
