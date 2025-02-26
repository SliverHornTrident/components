//go:build tencent && sec

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

var TencentSec = new(tencentSec)

type tencentSec struct{}

func (c *tencentSec) Name() string {
	return "[components][core][tencent][sec]"
}

func (c *tencentSec) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentSec", &component.TencentSecConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.TencentSecConfig.Tencent = component.TencentConfig
	if component.TencentSecConfig.SecretId == "" {
		component.TencentSecConfig.SecretId = component.TencentConfig.SecretId
	}
	if component.TencentSecConfig.SecretKey == "" {
		component.TencentSecConfig.SecretKey = component.TencentConfig.SecretKey
	}
	return nil
}

func (c *tencentSec) IsPanic() bool {
	return false
}

func (c *tencentSec) ConfigName() string {
	return strings.Join([]string{"tencent", "sec", gin.Mode(), "yaml"}, ".")
}

func (c *tencentSec) Initialization(ctx context.Context) error {
	sec, err := tencent.NewSec(component.TencentSecConfig)
	if err != nil {
		return err
	}
	component.TencentSec = sec
	return nil
}
