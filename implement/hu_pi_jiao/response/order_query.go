//go:build hu_pi_jiao

package response

type OrderQuery struct {
	Hash    string         `json:"hash"`
	ErrMsg  string         `json:"errmsg"`
	ErrCode int            `json:"errcode"`
	Data    OrderQueryData `json:"data"`
}

type OrderQueryData struct {
	Status      string `json:"status"` // OD(支付成功)，WP(待支付),CD(已取消)
	OpenOrderId string `json:"open_order_id"`
}
