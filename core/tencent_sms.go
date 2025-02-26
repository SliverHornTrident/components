//go:build tencent && sms

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/tencent"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var TencentSms = new(tencentSms)

type tencentSms struct{}

func (c *tencentSms) Name() string {
	return "[components][core][tencent][sms]"
}

func (c *tencentSms) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentSms", &component.TencentSmsConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.TencentSmsConfig.Tencent = component.TencentConfig
	if component.TencentSmsConfig.SecretId == "" {
		component.TencentSmsConfig.SecretId = component.TencentConfig.SecretId
	}
	if component.TencentSmsConfig.SecretKey == "" {
		component.TencentSmsConfig.SecretKey = component.TencentConfig.SecretKey
	}
	return nil
}

func (c *tencentSms) IsPanic() bool {
	return false
}

func (c *tencentSms) ConfigName() string {
	return strings.Join([]string{"tencent", "sms", gin.Mode(), "yaml"}, ".")
}

func (c *tencentSms) Initialization(ctx context.Context) error {
	sms, err := tencent.NewSms(component.TencentSmsConfig)
	if err != nil {
		return err
	}
	component.TencentSms = sms
	return nil
}
