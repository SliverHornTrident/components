//go:build aws

package config

type Aws struct {
	SecretID  string `json:"SecretID" yaml:"SecretID" mapstructure:"SecretID"`    // 密钥ID
	SecretKey string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 密钥Key
}
