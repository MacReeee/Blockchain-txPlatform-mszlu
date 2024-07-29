package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type MysqlConfig struct {
	DataSource string
}

type CqptchaConf struct {
	Vid string
	Key string
}

type Config struct {
	zrpc.RpcServerConf
	Mysql      MysqlConfig
	CacheRedis cache.CacheConf
	Captcha    CqptchaConf
}
