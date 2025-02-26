//go:build email

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Email = new(_email)

type _email struct{}

func (c *_email) Name() string {
	return "[components][core][email]"
}

func (c *_email) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Email", &component.EmailConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_email) IsPanic() bool {
	return false
}

func (c *_email) ConfigName() string {
	return strings.Join([]string{"email", gin.Mode(), "yaml"}, ".")
}

func (c *_email) Initialization(ctx context.Context) error {
	dialer := gomail.NewDialer(component.EmailConfig.Host, component.EmailConfig.Port, component.EmailConfig.Username, component.EmailConfig.Password)
	if component.EmailConfig.TLS {
		dialer.SSL = true
	}
	component.Email = dialer
	return nil
}
