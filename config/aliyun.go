//go:build aliyun

package config

type Aliyun struct {
	AccessKeyId     string `json:"AccessKeyId" yaml:"AccessKeyId" mapstructure:"AccessKeyId"`             // 访问密钥Id
	AccessKeySecret string `json:"AccessKeySecret" yaml:"AccessKeySecret" mapstructure:"AccessKeySecret"` // 访问密钥Secret
}
