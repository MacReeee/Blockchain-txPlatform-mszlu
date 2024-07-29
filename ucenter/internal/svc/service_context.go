package svc

import (
	"common/msdb"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"ucenter/internal/config"
	"ucenter/internal/database"
)

type ServiceContext struct {
	Config config.Config
	Cache  cache.Cache
	Db     *msdb.MsDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisCache := cache.New(c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config: c,
		Cache:  redisCache,
		Db:     database.ConnMysql(c.Mysql.DataSource),
	}
}
