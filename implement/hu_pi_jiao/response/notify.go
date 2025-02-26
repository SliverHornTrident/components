//go:build hu_pi_jiao

package response

type Notify struct {
	Hash          string  `json:"hash"`           // (32)签名=>未知=>请参考支付时签名算法)
	Time          string  `json:"time"`           // (16)时间戳=>未知=>无
	Appid         string  `json:"appid"`          // (32)支付渠道ID=>未知=>无
	Status        string  `json:"status"`         // (2)订单状态=>未知=>目前固定值为：OD
	Attach        string  `json:"attach"`         // (128)备注=>未知=>当传入此参数时才会有返回
	Plugins       string  `json:"plugins"`        // (128)插件ID=>未知=>当传入此参数时才会有返回
	NonceStr      string  `json:"nonce_str"`      // (16)随机字符串=>未知=>无
	OrderTitle    string  `json:"order_title"`    // (32)订单标题=>未知=>无
	OpenOrderId   string  `json:"open_order_id"`  // (32)虎皮椒内部订单号=>未知=>无
	TradeOrderId  string  `json:"trade_order_id"` // (32)商户订单号=>未知=>支付时请求的商户订单号
	TransactionId string  `json:"transaction_id"` // (32)交易号=>未知=>支付平台内部订单号
	TotalFee      float64 `json:"total_fee"`      // decimal(18,2)订单支付金额=>未知=>订单支付金额
}
