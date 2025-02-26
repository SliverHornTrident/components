//go:build tencent

package config

import "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

type Tencent struct {
	SecretId  string `json:"SecretId" yaml:"SecretId" mapstructure:"SecretId"`    // 访问密钥 Id
	SecretKey string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 访问密钥 Secret
}

func (r Tencent) NewCredential() *common.Credential {
	return common.NewCredential(r.SecretId, r.SecretKey)
}
