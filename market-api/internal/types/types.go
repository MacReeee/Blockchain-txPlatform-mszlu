package types

type RateRequest struct {
	Unit string `path:"unit" json:"unit"`
	Ip   string `json:"ip,optional"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}

type MarketReq struct {
	Ip         string `json:"ip,optional"`
	Symbol     string `json:"symbol,optional" form:"symbol,optional" path:"symbol,optional"`
	Unit       string `json:"unit,optional" form:"unit,optional"`
	From       int64  `json:"from,optional" form:"from,optional"`
	To         int64  `json:"to,optional" form:"to,optional"`
	Resolution string `json:"resolution,optional" form:"resolution,optional"`
}

type CoinThumbResp struct {
	Symbol       string    `json:"symbol"`
	Open         float64   `json:"open"`
	High         float64   `json:"high"`
	Low          float64   `json:"low"`
	Close        float64   `json:"close"`
	Chg          float64   `json:"chg"`    //变化百分比
	Change       float64   `json:"change"` // 变化金额
	Volume       float64   `json:"volume"`
	Turnover     float64   `json:"turnover"`
	LastDayClose float64   `json:"lastDayClose"`
	UsdRate      float64   `json:"usdRate"`        // USDT汇率
	BaseUsdRate  float64   `json:"baseUsdRate"`    // 基础USDT汇率
	Zone         int       `json:"zone"`           // 交易区
	Trend        []float64 `json:"trend,optional"` //价格趋势
}

type ExchangeCoinResp struct {
	Id                 int64   `json:"id"`
	Symbol             string  `json:"symbol"`             // 交易对币种名称，格式：BTC/USDT
	BaseCoinScale      int64   `json:"baseCoinScale"`      // 基币小数精度
	BaseSymbol         string  `json:"baseSymbol"`         // 结算币种符号，如USDT
	CoinScale          int64   `json:"coinScale"`          // 交易币小数精度
	CoinSymbol         string  `json:"coinSymbol"`         // 交易币种符号
	Enable             int64   `json:"enable"`             // 状态，1：启用，2：禁止
	Fee                float64 `json:"fee"`                // 交易手续费
	Sort               int64   `json:"sort"`               // 排序，从小到大
	EnableMarketBuy    int64   `json:"enableMarketBuy"`    // 是否启用市价买
	EnableMarketSell   int64   `json:"enableMarketSell"`   // 是否启用市价卖
	MinSellPrice       float64 `json:"minSellPrice"`       // 最低挂单卖价
	Flag               int64   `json:"flag"`               // 标签位，用于推荐，排序等,默认为0，1表示推荐
	MaxTradingOrder    int64   `json:"maxTradingOrder"`    // 最大允许同时交易的订单数，0表示不限制
	MaxTradingTime     int64   `json:"maxTradingTime"`     // 委托超时自动下架时间，单位为秒，0表示不过期
	MinTurnover        float64 `json:"minTurnover"`        // 最小挂单成交额
	ClearTime          int64   `json:"clearTime"`          // 清盘时间
	EndTime            int64   `json:"endTime"`            // 结束时间
	Exchangeable       int64   `json:"exchangeable"`       //  是否可交易
	MaxBuyPrice        float64 `json:"maxBuyPrice"`        // 最高买单价
	MaxVolume          float64 `json:"maxVolume"`          // 最大下单量
	MinVolume          float64 `json:"minVolume"`          // 最小下单量
	PublishAmount      float64 `json:"publishAmount"`      //  活动发行数量
	PublishPrice       float64 `json:"publishPrice"`       //  分摊发行价格
	PublishType        int64   `json:"publishType"`        // 发行活动类型 1:无活动,2:抢购发行,3:分摊发行
	RobotType          int64   `json:"robotType"`          // 机器人类型
	StartTime          int64   `json:"startTime"`          // 开始时间
	Visible            int64   `json:"visible"`            //  前台可见状态
	Zone               int64   `json:"zone"`               // 交易区域
	CurrentTime        int64   `json:"currentTime"`        //当前毫秒值
	MarketEngineStatus int     `json:"marketEngineStatus"` //行情引擎状态（0：不可用，1：可用
	EngineStatus       int     `json:"engineStatus"`       //交易引擎状态（0：不可用，1：可用
	ExEngineStatus     int     `json:"exEngineStatus"`     //交易机器人状态（0：非运行中，1：运行中）
}

type Coin struct {
	Id                int     `json:"id" from:"id"`
	Name              string  `json:"name" from:"name"`
	CanAutoWithdraw   int     `json:"canAutoWithdraw" from:"canAutoWithdraw"`
	CanRecharge       int     `json:"canRecharge" from:"canRecharge"`
	CanTransfer       int     `json:"canTransfer" from:"canTransfer"`
	CanWithdraw       int     `json:"canWithdraw" from:"canWithdraw"`
	CnyRate           float64 `json:"cnyRate" from:"cnyRate"`
	EnableRpc         int     `json:"enableRpc" from:"enableRpc"`
	IsPlatformCoin    int     `json:"isPlatformCoin" from:"isPlatformCoin"`
	MaxTxFee          float64 `json:"maxTxFee" from:"maxTxFee"`
	MaxWithdrawAmount float64 `json:"maxWithdrawAmount" from:"maxWithdrawAmount"`
	MinTxFee          float64 `json:"minTxFee" from:"minTxFee"`
	MinWithdrawAmount float64 `json:"minWithdrawAmount" from:"minWithdrawAmount"`
	NameCn            string  `json:"nameCn" from:"nameCn"`
	Sort              int     `json:"sort" from:"sort"`
	Status            int     `json:"status" from:"status"`
	Unit              string  `json:"unit" from:"unit"`
	UsdRate           float64 `json:"usdRate" from:"usdRate"`
	WithdrawThreshold float64 `json:"withdrawThreshold" from:"withdrawThreshold"`
	HasLegal          int     `json:"hasLegal" from:"hasLegal"`
	ColdWalletAddress string  `json:"coldWalletAddress" from:"coldWalletAddress"`
	MinerFee          float64 `json:"minerFee" from:"minerFee"`
	WithdrawScale     int     `json:"withdrawScale" from:"withdrawScale"`
	AccountType       int     `json:"accountType" from:"accountType"`
	DepositAddress    string  `json:"depositAddress" from:"depositAddress"`
	Infolink          string  `json:"infolink" from:"infolink"`
	Information       string  `json:"information" from:"information"`
	MinRechargeAmount float64 `json:"minRechargeAmount" from:"minRechargeAmount"`
}

//	message History {
//	 int64 time = 1;
//	 double open = 2;
//	 double close = 3;
//	 double high = 4;
//	 double low = 5;
//	 double volume = 6;
//	}
type HistoryKline struct {
	List [][]any
}
