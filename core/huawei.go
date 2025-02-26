//go:build huawei

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Huawei = new(_huawei)

type _huawei struct{}

func (c *_huawei) Name() string {
	return "[components][core][huawei]"
}

func (c *_huawei) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Huawei", &component.HuaweiConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_huawei) IsPanic() bool {
	return false
}

func (c *_huawei) ConfigName() string {
	return strings.Join([]string{"huawei", gin.Mode(), "yaml"}, ".")
}

func (c *_huawei) Initialization(ctx context.Context) error {
	return nil
}
