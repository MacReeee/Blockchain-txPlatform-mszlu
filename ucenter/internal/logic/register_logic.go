package logic

import (
	"context"
	"grpc-common/ucenter/types/register"

	"github.com/zeromicro/go-zero/core/logx"
	"ucenter/internal/svc"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (*register.RegRes, error) {
	logx.Info("ucenter rpc register by phone called...")
	return &register.RegRes{}, nil
}
