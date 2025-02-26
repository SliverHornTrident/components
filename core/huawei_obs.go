//go:build huawei && obs

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var HuaweiObs = new(huaweiObs)

type huaweiObs struct{}

func (c *huaweiObs) Name() string {
	return "[components][core][huawei][obs]"
}

func (c *huaweiObs) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("HuaweiObs", &component.HuaweiObsConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.HuaweiObsConfig.Huawei = component.HuaweiConfig
	return nil
}

func (c *huaweiObs) IsPanic() bool {
	return false
}

func (c *huaweiObs) ConfigName() string {
	return strings.Join([]string{"huawei", "obs", gin.Mode(), "yaml"}, ".")
}

func (c *huaweiObs) Initialization(ctx context.Context) error {
	return nil
}
