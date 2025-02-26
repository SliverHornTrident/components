//go:build tencent && ses

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

var TencentSes = new(tencentSes)

type tencentSes struct{}

func (c *tencentSes) Name() string {
	return "[components][core][tencent][ses]"
}

func (c *tencentSes) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("TencentSes", &component.TencentSesConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *tencentSes) IsPanic() bool {
	return false
}

func (c *tencentSes) ConfigName() string {
	return strings.Join([]string{"tencent", "ses", gin.Mode(), "yaml"}, ".")
}

func (c *tencentSes) Initialization(ctx context.Context) error {
	ses := tencent.NewSes(component.TencentSesConfig)
	component.TencentSes = ses
	return nil
}
