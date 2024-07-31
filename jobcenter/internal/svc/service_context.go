package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/ucclient"
	"ucenter-api/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RegisterRpc ucclient.Register
	UCLoginRpc  ucclient.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RegisterRpc: ucclient.NewRegister(zrpc.MustNewClient(c.UCenterRpc)),
		UCLoginRpc:  ucclient.NewLogin(zrpc.MustNewClient(c.UCenterRpc)),
	}
}