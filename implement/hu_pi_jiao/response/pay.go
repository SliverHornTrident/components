//go:build hu_pi_jiao

package response

type Pay struct {
	Url       string `json:"url"`        // (155)请求url(手机端专用，PC端已停用)=>未知=>只需跳转此参数即可，系统会自动判断是微信端还是手机端，自动返回return_url，不能先显示“url_qrcode”二维码，再跳转“url”链接
	Hash      string `json:"hash"`       // (32)签名=>未知=>请参考支付时签名算法)
	ErrMsg    string `json:"errmsg"`     // (8)错误信息=>未知=>错误信息具体值
	OderId    string `json:"oderid"`     // (unknown)订单id=>未知=>订单id
	UrlQrcode string `json:"url_qrcode"` // (155)二维码地址(PC端使用)=>未知=>PC端可将该参数展示出来进行扫码支付，不用再转二维码，需自己处理跳转
	ErrCode   int    `json:"errcode"`    // (unknown)错误码=>未知=>无
}
