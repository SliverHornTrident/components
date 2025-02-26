//go:build tencent

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Tencent = new(_tencent)

type _tencent struct{}

func (c *_tencent) Name() string {
	return "[components][core][tencent]"
}

func (c *_tencent) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Tencent", &component.TencentConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_tencent) IsPanic() bool {
	return false
}

func (c *_tencent) ConfigName() string {
	return strings.Join([]string{"tencent", gin.Mode(), "yaml"}, ".")
}

func (c *_tencent) Initialization(ctx context.Context) error {
	return nil
}
