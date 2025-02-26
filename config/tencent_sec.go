//go:build tencent && sec

package config

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type TencentSec struct {
	Region       string  `json:"Region" yaml:"Region" mapstructure:"Region"`                          // 区域
	Endpoint     string  `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`                    // 端点
	SecretId     string  `json:"SecretId" yaml:"SecretId" mapstructure:"SecretId"`                    // 访问密钥 Id
	SecretKey    string  `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"`                 // 访问密钥 Secret
	AppSecretKey string  `json:"AppSecretKey" yaml:"AppSecretKey" mapstructure:"AppSecretKey"`        // 验证码应用密钥
	CaptchaType  uint64  `json:"CaptchaType,string" yaml:"CaptchaType" mapstructure:"CaptchaType"`    // 验证码类型
	CaptchaAppId uint64  `json:"CaptchaAppId,string" yaml:"CaptchaAppId" mapstructure:"CaptchaAppId"` // 验证码应用ID
	Tencent      Tencent `json:"Tencent" yaml:"-" mapstructure:"-"`
}

func (c *TencentSec) NewCredential() *common.Credential {
	if c.SecretId != "" {
		c.Tencent.SecretId = c.SecretId
	}
	if c.SecretKey != "" {
		c.Tencent.SecretKey = c.SecretKey
	}
	return c.Tencent.NewCredential()
}

func (c *TencentSec) ClientProfile() *profile.ClientProfile {
	return &profile.ClientProfile{
		Language: "zh-CN",
		HttpProfile: &profile.HttpProfile{
			Endpoint: c.Endpoint,
		},
	}
}
