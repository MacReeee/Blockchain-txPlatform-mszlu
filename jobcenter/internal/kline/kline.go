package kline

import (
	"common/tools"
	"encoding/json"
	"log"
	"sync"
	"time"
)

type OkxConfig struct {
	ApiKey    string
	SecretKey string
	Pass      string
	Host      string
}

var secretKey = "8F61D46421A57599A5CAFD26D17D5BE0"

type OkxResult struct {
	Code string     `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}

type Kline struct {
	wg sync.WaitGroup
}

func (k *Kline) Do(period string) {
	k.wg.Add(2)
	k.getKlineData("BTC-USDT", period)
	k.getKlineData("ETH-USDT", period)
	k.wg.Wait()
}

func (k *Kline) getKlineData(instId string, period string) {
	//获取数据
	api := "https://www.okx.com/api/v5/market/candles?instId=" + instId + "&bar=" + period
	timestamp := tools.ISO(time.Now())
	sign := tools.ComputeHmacSha256(timestamp+"GET/api/v5/market/candles?instId="+instId+"&bar="+period, secretKey)
	header := make(map[string]string)
	header["OK-ACCESS-KEY"] = "d5a748c6-214d-4fae-bef3-d32368ecbbe8"
	header["OK-ACCESS-SIGN"] = sign
	header["OK-ACCESS-TIMESTAMP"] = timestamp
	header["OK-ACCESS-PASSPHRASE"] = "MacReeee1@github.com"
	resp, err := tools.GetWithHeader(api, header, "http://localhost:7890")
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
	log.Println("============获取到的k线数据==============")
	log.Println("instId: ", instId, "period: ", period)
	log.Println("result kline data: ", string(resp))
	log.Println("============End==============")
	k.wg.Done()
}

func NewKline() *Kline {
	return &Kline{}
}
