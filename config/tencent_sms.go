//go:build tencent && sms

package config

import (
	"crypto/rand"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"math/big"
)

type TencentSms struct {
	AppId      string  `json:"AppId" yaml:"AppId" mapstructure:"AppId"`                // 短信应用 ID
	SignName   string  `json:"SignName" yaml:"SignName" mapstructure:"SignName"`       // 短信签名
	TemplateId string  `json:"TemplateId" yaml:"TemplateId" mapstructure:"TemplateId"` // 短信模板 ID
	Region     string  `json:"Region" yaml:"Region" mapstructure:"Region"`             // 地域: ap-beijing(华北地区=>北京) || ap-guangzhou(华南地区=>广州) || ap-nanjing(华东地区=>南京)
	Endpoint   string  `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`       // 短信服务域名
	SecretId   string  `json:"SecretId" yaml:"SecretId" mapstructure:"SecretId"`       // 访问密钥 Id
	SecretKey  string  `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"`    // 访问密钥 Secret
	Tencent    Tencent `json:"Tencent" yaml:"-" mapstructure:"-"`
}

func (c *TencentSms) NewCredential() *common.Credential {
	if c.SecretId != "" {
		c.Tencent.SecretId = c.SecretId
	}
	if c.SecretKey != "" {
		c.Tencent.SecretKey = c.SecretKey
	}
	return c.Tencent.NewCredential()
}

func (c *TencentSms) TemplateParam(captcha int, phone ...*string) []*string {
	length := len(phone)
	codes := make([]*string, 0, length)
	for i := 0; i < length; i++ {
		var code string
		for j := 0; j < captcha; j++ {
			result, _ := rand.Int(rand.Reader, big.NewInt(10))
			code += result.String()
		}
		codes = append(codes, &code)
	}
	return codes
}
