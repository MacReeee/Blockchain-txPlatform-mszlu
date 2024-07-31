package kline

import (
	"common/tools"
	"encoding/json"
	"jobcenter/internal/database"
	"jobcenter/internal/domain"
	"log"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
	Proxy     string
}

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg          sync.WaitGroup
	c           OkxConfig
	klineDomain *domain.KlineDomain
	queueDomain *domain.QueueDomain
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	go k.getKlineData("BTC-USDT", "BTC/USDT", period)
	go k.getKlineData("ETH-USDT", "ETH/USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId, symbol, period string) {
	//获取数据
	api := k.c.Host + "/api/v5/market/candles?instId=" + instId + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET/api/v5/market/candles?instId="+instId+"&bar="+period, k.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = k.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = k.c.Pass
	resp, err := tools.GetWithHeader(api, header, k.c.Proxy)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	res := &OkxResult{}
	err = json.Unmarshal(resp, res)
	if err != nil {
		log.Println(err)
		k.wg.Done()
		return
	}
	log.Println("============执行存储mongo==============")
	if res.Code == "0" {
		k.klineDomain.SaveBatch(res.Data, symbol, period)
		if period == "1m" {
			if len(res.Data) > 0 {
				k.queueDomain.Send1mKline(res.Data[0], symbol)
			}
		}
	}
	log.Println("============End==============")
	k.wg.Done()
}

func NewKline(c OkxConfig, mongoCLi *database.MongoClient, kafkaCli *database.KafkaClient) *Kline {
	return &Kline{
		c:           c,
		klineDomain: domain.NewKlineDomain(mongoCLi),
		queueDomain: domain.NewQueueDomain(kafkaCli),
	}
}
