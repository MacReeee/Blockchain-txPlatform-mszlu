// Code generated by goctl. DO NOT EDIT.
// Source: rate

package server

import (
	"context"
	"grpc-common/ucenter/types/register"

	"ucenter/internal/logic"
	"ucenter/internal/svc"
)

type RegisterServer struct {
	svcCtx *svc.ServiceContext
	register.UnimplementedRegisterServer
}

func NewRegisterServer(svcCtx *svc.ServiceContext) *RegisterServer {
	return &RegisterServer{
		svcCtx: svcCtx,
	}
}

func (s *RegisterServer) RegisterByPhone(ctx context.Context, in *register.RegReq) (*register.RegRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.RegisterByPhone(in)
}

func (s *RegisterServer) SendCode(ctx context.Context, in *register.CodeReq) (*register.NoRes, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.SendCode(in)
}
