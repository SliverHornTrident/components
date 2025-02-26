//go:build tencent && cdn

package config

import "time"

type TencentCdn []TencentCdnChildren

type TencentCdnChildren struct {
	Key          string        `json:"Key" yaml:"Key" mapstructure:"Key"`                            // 密钥
	Mode         string        `json:"Mode" yaml:"Mode" mapstructure:"Mode"`                         // 鉴权模式
	Domain       string        `json:"Domain" yaml:"Domain" mapstructure:"Domain"`                   // cdn加速域名
	BackupKey    string        `json:"BackupKey" yaml:"BackupKey" mapstructure:"BackupKey"`          // 备用密钥
	Algorithm    string        `json:"Algorithm" yaml:"Algorithm" mapstructure:"Algorithm"`          // 鉴权算法
	SignatureKey string        `json:"SignatureKey" yaml:"SignatureKey" mapstructure:"SignatureKey"` // 签名密钥返回参数名
	TimestampKey string        `json:"TimestampKey" yaml:"TimestampKey" mapstructure:"TimestampKey"` // 时间戳返回参数名
	Expire       time.Duration `json:"Expire" yaml:"Expire" mapstructure:"Expire"`                   // 过期时间
	ExpireBase   int           `json:"ExpireBase" yaml:"ExpireBase" mapstructure:"ExpireBase"`       // Unix进制(10/16)
}
