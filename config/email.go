//go:build email

package config

type Email struct {
	Host     string `json:"Host" yaml:"Host" mapstructure:"Host"`
	Username string `json:"Username" yaml:"Username" mapstructure:"Username"`
	Password string `json:"Password" yaml:"Password" mapstructure:"Password"`
	Port     int    `json:"Port" yaml:"Port" mapstructure:"Port"`
	TLS      bool   `json:"TLS" yaml:"TLS" mapstructure:"TLS"`
}
