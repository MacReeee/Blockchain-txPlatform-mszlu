package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type AuthConfig struct {
	AccessSecret string
	AccessExpire int64
}

type Config struct {
	rest.RestConf
	UCenterRpc zrpc.RpcClientConf
	JWT        AuthConfig
}
