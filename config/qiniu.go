//go:build qiniu

package config

type Qiniu struct {
	AccessKey string `json:"AccessKey" yaml:"AccessKey" mapstructure:"AccessKey"` // AccessKey 秘钥AK
	SecretKey string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // SecretKey 秘钥SK
}
