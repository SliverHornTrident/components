//go:build hu_pi_jiao

package config

type HuPiJiao []HuPiJiaoChildren

type HuPiJiaoChildren struct {
	Name              string `json:"Name" yaml:"Name" mapstructure:"Name"`
	AppId             string `json:"AppId" yaml:"AppId" mapstructure:"AppId"`
	Gateway           string `json:"Gateway" yaml:"Gateway" mapstructure:"Gateway"`
	AppSecret         string `json:"AppSecret" yaml:"AppSecret" mapstructure:"AppSecret"`
	ReturnUrl         string `json:"ReturnUrl" yaml:"ReturnUrl" mapstructure:"ReturnUrl"`                         // 同步通知页面 HTTP/HTTPS开头字符串
	NotifyUrl         string `json:"NotifyUrl" yaml:"NotifyUrl" mapstructure:"NotifyUrl"`                         // 虎皮椒服务器主动通知商户服务器里指定的页面http/https路径。
	PayGateway        string `json:"PayGateway" yaml:"PayGateway" mapstructure:"PayGateway"`                      // 支付网关
	CallbackUrl       string `json:"CallbackUrl" yaml:"CallbackUrl" mapstructure:"CallbackUrl"`                   // 用户取消支付后，我们可能引导用户跳转到这个网址上重新进行支付
	OrderQueryGateway string `json:"OrderQueryGateway" yaml:"OrderQueryGateway" mapstructure:"OrderQueryGateway"` // 订单查询网关
}
