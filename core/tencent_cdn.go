//go:build tencent && cdn

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

var TencentCdn = new(tencentCdn)

type tencentCdn struct{}

func (c *tencentCdn) Name() string {
	return "[components][core][tencent][cdn]"
}

func (c *tencentCdn) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentCdn", &component.TencentCdnConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal config failed!")
	}
	return nil
}

func (c *tencentCdn) IsPanic() bool {
	return false
}

func (c *tencentCdn) ConfigName() string {
	return strings.Join([]string{"tencent", "cdn", gin.Mode(), "yaml"}, ".")
}

func (c *tencentCdn) Initialization(ctx context.Context) error {
	length := len(component.TencentCdnConfig)
	component.TencentCdn = make(map[string]*tencent.Cdn, length)
	for i := 0; i < length; i++ {
		cdn := tencent.NewCdn(component.TencentCdnConfig[i])
		component.TencentCdn[component.TencentCdnConfig[i].Domain] = cdn
	}
	return nil
}
