//go:build jwt

package config

import "time"

type Jwt struct {
	Issuer     string        `json:"Issuer" yaml:"Issuer" mapstructure:"Issuer"`             // 签发者
	BufferAt   time.Duration `json:"BufferAt" yaml:"BufferAt" mapstructure:"BufferAt"`       // 缓冲时间
	ExpiresAt  time.Duration `json:"ExpiresAt" yaml:"ExpiresAt" mapstructure:"ExpiresAt"`    // 过期时间
	SigningKey string        `json:"SigningKey" yaml:"SigningKey" mapstructure:"SigningKey"` // jwt签名
}
