//go:build hu_pi_jiao

package request

type Pay struct {
	Hash         string  `json:"hash,omitempty"`             // (32)签名=>必填=>无
	Type         string  `json:"type,omitempty"`             // (32)支付通道类型=>未知=>微信H5支付请填"WAP"，微信小程序支付请填"JSAPI" ，请参考小程序demo对接小程序支付，微信内支付请勿填写"JSAPI"，支付网关为：https://api.xunhupay.com 跳转小程序APPID：wx2574b5c5ee8da56b，其他支付网关跳转小程序APPID：wx402faa5bd5eda155，支付宝不需要此参数
	Appid        string  `json:"appid,omitempty"`            // (32)APP ID=>必填=>填写虎皮椒的APPID，不是小程序APPID
	Title        string  `json:"title,omitempty"`            // (128)订单标题=>必填=>商户订单标题（不能超过127个字符或者63个汉字，请注意控制下长度）
	WapUrl       string  `json:"wap_url,omitempty"`          // (128)网站域名=>网站域名，H5支付通道请填你网站域名，小程序支付请填支付网关(例如：https://api.dpweixin.com)。支付宝不需要此参数
	Attach       string  `json:"attach,omitempty"`           // (∞)备注	text=>可选=>备注字段，可以传入一些备注数据，回调时原样返回
	Plugins      string  `json:"plugins,omitempty"`          // (128)名称=>可选=>用于识别对接程序或作者
	WapName      string  `json:"wap_name,omitempty"`         // (128)网站名称=>未知=>店铺名称或网站域名，长度32或以内，H5支付通道请求必填。支付宝不需要此参数
	Version      string  `json:"version,omitempty"`          // (24)API 版本号=>必填=>目前为1.1
	NonceStr     string  `json:"nonce_str,omitempty"`        // (32)随机值=>必填=>作用：1.避免服务器页面缓存，2.防止安全密钥被猜测出来
	NotifyUrl    string  `json:"notify_url,omitempty"`       // (128)通知回调网址=>必填。	(用户支付成功后，我们服务器会主动发送一个post消息到这个网址(注意：当前接口内，SESSION内容无效，手机端不支持中文域名))
	ReturnUrl    string  `json:"return_url,omitempty"`       // (128)跳转网址=>可选=>用户支付成功后，我们会让用户浏览器自动跳转到这个网址
	CallbackUrl  string  `json:"callback_url,omitempty"`     // (128)商品网址=>可选=>用户取消支付后，我们可能引导用户跳转到这个网址上重新进行支付
	TradeOrderId string  `json:"trade_order_id,omitempty"`   // (32)商户订单号=>必填=>请确保在当前网站内是唯一订单号，只支持数字，大小写英文以及部分特殊字符：!#$'()*+,/:;=?@-._~%
	Time         int64   `json:"time,string,omitempty"`      // (11)当前时间戳=>必填=>PHP示例：time()
	TotalFee     float64 `json:"total_fee,string,omitempty"` // decimal(18,2)订单金额(元)>必填=>位为人民币 元，没小数位不用强制保留2位小数
}
