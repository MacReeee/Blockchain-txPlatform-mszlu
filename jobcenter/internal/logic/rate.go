package logic

import (
	"common/tools"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"log"
	"sync"
	"time"
)

type Rate struct {
	wg sync.WaitGroup
	c  OkxConfig
	ch cache.Cache
}

func (r *Rate) Do() {
	//要获取 人民币对美元的汇率
	r.wg.Add(1)
	go r.CnyUsdRate()
	r.wg.Wait()
}

type OkxExchangeRateResult struct {
	Code string         `json:"code"`
	Msg  string         `json:"msg"`
	Data []ExchangeRate `json:"data"`
}
type ExchangeRate struct {
	UsdCny string `json:"usdCny"`
}

func (r *Rate) CnyUsdRate() {
	//请求接口 获取到最新的汇率 存入redis即可
	//发起http请求 获取数据
	api := r.c.Host + "/api/v5/market/exchange-rate"
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET"+"/api/v5/market/exchange-rate", r.c.SecretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = r.c.ApiKey
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = r.c.Pass
	resp, err := tools.GetWithHeader(api, header, r.c.Proxy)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	var result = &OkxExchangeRateResult{}
	err = json.Unmarshal(resp, result)
	if err != nil {
		log.Println(err)
		r.wg.Done()
		return
	}
	cny := result.Data[0].UsdCny
	//存入redis
	r.ch.Set("USDT::CNY::RATE", cny)
	r.wg.Done()
}

func NewRate(c OkxConfig, cache2 cache.Cache) *Rate {
	return &Rate{
		c:  c,
		ch: cache2,
	}
}
