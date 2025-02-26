//go:build huawei

package config

type Huawei struct {
	AccessKey string `json:"AccessKey" yaml:"AccessKey" mapstructure:"AccessKey"` // 访问密钥
	SecretKey string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 密钥
}
