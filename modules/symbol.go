package modules

/*

eg:
https://stock.xueqiu.com/v5/stock/realtime/quotec.json?symbol=SH600795
https://stock.xueqiu.com/v5/stock/realtime/quotec.json?symbol=SH600795,SH000001

{
  "data": [
    {
      "symbol": "SH600795",
      "current": 5.25,
      "percent": 0.57,
      "chg": 0.03,
      "timestamp": 1729666800000,
      "volume": 155462961,
      "amount": 812128545.36,
      "market_capital": 93637000181,
      "float_market_capital": 93637000181,
      "turnover_rate": 0.87,
      "amplitude": 2.3,
      "open": 5.26,
      "last_close": 5.22,
      "high": 5.28,
      "low": 5.16,
      "avg_price": 5.22393591461313,
      "trade_volume": null,
      "side": null,
      "is_trade": false,
      "level": 2,
      "trade_session": null,
      "trade_type": null,
      "current_year_percent": 27.69,
      "trade_unique_id": null,
      "type": 11,
      "bid_appl_seq_num": null,
      "offer_appl_seq_num": null,
      "volume_ext": null,
      "traded_amount_ext": null,
      "trade_type_v2": null,
      "yield_to_maturity": null
    }
  ],
  "error_code": 0,
  "error_description": null
}

*/

type OriginData struct {
	Data             []SymbolData `json:"data"`       //元数据
	ErrorCode        int          `json:"error_code"` //错误码
	ErrorDescription string       `json:"错误描述"`       //错误描述
}

type SymbolData struct {
	Symbol  string  `json:"symbol"`  //编号
	Current float64 `json:"current"` // 当前价
	Percent float64 `json:"percent"` // 涨跌幅
	Chg     float64 `json:"chg"`     // 涨跌额
	//Timestamp          int64    `json:"timestamp"`            // 时间戳（毫秒）
	//Volume             uint64   `json:"volume"`               // 成交量
	//Amount             float64  `json:"amount"`               // 成交金额
	//MarketCapital      uint64   `json:"market_capital"`       // 总市值
	//FloatMarketCapital uint64   `json:"float_market_capital"` // 浮动市值
	//TurnoverRate       float64  `json:"turnover_rate"`        // 换手率
	//Amplitude          float64  `json:"amplitude"`            // 振幅
	Open float64 `json:"open"` // 开盘价
	//LastClose          float64  `json:"last_close"`           // 最近收盘价
	High float64 `json:"high"` // 高点
	Low  float64 `json:"low"`  // 低点
	//AveragePrice       float64  `json:"avg_price"`            // 均价
	//TradeVolume        *uint64  `json:"trade_volume"`         //  成交量
	//Side               *string  `json:"side"`                 // 多空
	//IsTrade            bool     `json:"is_trade"`             // 是否交易
	//Level              int      `json:"level"`                // 市场等级
	//TradeSession       *int     `json:"trade_session"`        // 交易会序号
	//TradeType          *string  `json:"trade_type"`           // 交易类型
	//CurrentYearPercent float64  `json:"current_year_percent"` // 当年百分比
	//TradeUniqueId      *string  `json:"trade_unique_id"`      // 唯一交易ID
	//Type               int      `json:"type"`                 // 证券类别代码
	//BidApplSeqNum      *uint64  `json:"bid_appl_seq_num"`     // 委买方申请序列号
	//OfferApplSeqNum    *uint64  `json:"offer_appl_seq_num"`   // 委卖方申请序列号
	//VolumeExt          *uint64  `json:"volume_ext"`           // 增减仓数量
	//TradedAmountExt    *float64 `json:"traded_amount_ext"`    // 增减金额
	//TradeTypeV2        *string  `json:"trade_type_v2"`        // V2版本的交易类型
	//YieldToMaturity    *float64 `json:"yield_to_maturity"`    // 到期收益率
}
