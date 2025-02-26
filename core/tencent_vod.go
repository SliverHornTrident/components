//go:build tencent && vod

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

var TencentVod = new(tencentVod)

type tencentVod struct{}

func (c *tencentVod) Name() string {
	return "[components][core][tencent][vod]"
}

func (c *tencentVod) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentVod", &component.TencentVodConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.TencentVodConfig.Tencent = component.TencentConfig
	if component.TencentVodConfig.SecretId == "" {
		component.TencentVodConfig.SecretId = component.TencentConfig.SecretId
	}
	if component.TencentVodConfig.SecretKey == "" {
		component.TencentVodConfig.SecretKey = component.TencentConfig.SecretKey
	}
	return nil
}

func (c *tencentVod) IsPanic() bool {
	return false
}

func (c *tencentVod) ConfigName() string {
	return strings.Join([]string{"tencent", "vod", gin.Mode(), "yaml"}, ".")
}

func (c *tencentVod) Initialization(ctx context.Context) error {
	vod := tencent.NewVod(component.TencentVodConfig)
	component.TencentVod = vod
	return nil
}
