//go:build tencent && cos

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

var TencentCos = new(tencentCos)

type tencentCos struct{}

func (c *tencentCos) Name() string {
	return "[components][core][tencent][cos]"
}

func (c *tencentCos) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentCos", &component.TencentCosConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.TencentCosConfig.Tencent = component.TencentConfig
	if component.TencentCosConfig.SecretId == "" {
		component.TencentCosConfig.SecretId = component.TencentConfig.SecretId
	}
	if component.TencentCosConfig.SecretKey == "" {
		component.TencentCosConfig.SecretKey = component.TencentConfig.SecretKey
	}
	return nil
}

func (c *tencentCos) IsPanic() bool {
	return false
}

func (c *tencentCos) ConfigName() string {
	return strings.Join([]string{"tencent", "cos", gin.Mode(), "yaml"}, ".")
}

func (c *tencentCos) Initialization(ctx context.Context) error {
	cos, err := tencent.NewCos(component.TencentCosConfig)
	if err != nil {
		return err
	}
	component.TencentCos = cos
	return nil
}
