package svc

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"grpc-common/ucenter/ucclient"
	"jobcenter/internal/config"
	"jobcenter/internal/database"
)

type ServiceContext struct {
	Config         config.Config
	MongoClient    *database.MongoClient
	KafkaClient    *database.KafkaClient
	Cache          cache.Cache
	AssetRpc       ucclient.Asset
	BitCoinAddress string
}

func NewServiceContext(c config.Config) *ServiceContext {
	client := database.NewKafkaClient(c.Kafka)
	client.StartWrite()
	redisCache := cache.New(
		c.CacheRedis,
		nil,
		cache.NewStat("mscoin"),
		nil,
		func(o *cache.Options) {})
	return &ServiceContext{
		Config:         c,
		MongoClient:    database.ConnectMongo(c.Mongo),
		KafkaClient:    client,
		Cache:          redisCache,
		AssetRpc:       ucclient.NewAsset(zrpc.MustNewClient(c.UCenterRpc)),
		BitCoinAddress: c.Bitcoin.Address,
	}
}
