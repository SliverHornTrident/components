//go:build hu_pi_jiao

package request

type OrderQuery struct {
	Hash          string `json:"hash,omitempty"`            // (32)签名=>必填=>无
	Time          int64  `json:"time,string,omitempty"`     // 当前时间戳=>必填=>PHP示例：time()
	AppId         string `json:"appid,omitempty"`           // (32)平台分配商户号=>必填=>应用ID
	NonceStr      string `json:"nonce_str,omitempty"`       // (32)随机值=>必填=>作用：1.避免服务器页面缓存，2.防止安全密钥被猜测出来
	OpenOrderId   string `json:"open_order_id,omitempty"`   // (32)虎皮椒内部订单号=>必填=>out_trade_order，open_order_id 二选一。在支付时，或支付成功时会返回此数据给商户网站y
	OutTradeOrder string `json:"out_trade_order,omitempty"` // (32)商户网站订单号=>必填=>out_trade_order，open_order_id 二选一。请确保在您的网站内是唯一订单号
}
