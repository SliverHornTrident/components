//go:build tencent && ses

package config

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/smtp"
	"net/url"
	"strings"
)

type TencentSes struct {
	Port     uint   `json:"Port" yaml:"Port" mapstructure:"Port"`
	Host     string `json:"Host" yaml:"Host" mapstructure:"Host"`
	Username string `json:"Username" yaml:"Username" mapstructure:"Username"`
	Password string `json:"Password" yaml:"Password" mapstructure:"Password"`
	TLS      bool   `json:"TLS" yaml:"TLS" mapstructure:"TLS"`
}

func (c *TencentSes) Header(subject string, tos []string) url.Values {
	values := make(url.Values, 4)
	values.Add("To", strings.Join(tos, ","))
	values.Add("From", c.Username)
	values.Add("Subject", subject)
	values.Add("Content-Type", "text/html; charset=UTF-8")
	return values
}

func (c *TencentSes) Auth() smtp.Auth {
	return smtp.PlainAuth("", c.Username, c.Password, c.Host)
}

func (c *TencentSes) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *TencentSes) Captcha(captcha int, email ...string) map[string]string {
	length := len(email)
	codes := make(map[string]string, length)
	for i := 0; i < length; i++ {
		var code string
		for j := 0; j < captcha; j++ {
			result, _ := rand.Int(rand.Reader, big.NewInt(10))
			code += result.String()
		}
		codes[email[i]] = code
	}
	return codes
}
