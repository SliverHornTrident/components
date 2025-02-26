//go:build sonyflake

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var Sonyflake = new(_sonyflake)

type _sonyflake struct{}

func (c *_sonyflake) Name() string {
	return "[components][core][snowflake]"
}

func (c *_sonyflake) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Sonyflake", &component.SonyflakeConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *_sonyflake) IsPanic() bool {
	return false
}

func (c *_sonyflake) ConfigName() string {
	return strings.Join([]string{"snowflake", gin.Mode(), "yaml"}, ".")
}

func (c *_sonyflake) Initialization(ctx context.Context) error {
	startTime, err := time.ParseInLocation(time.DateTime, component.SonyflakeConfig.Start, time.Local)
	if err != nil {
		return errors.Wrap(err, "[components][core][sonyflake] 时间格式错误!")
	}
	settings := sonyflake.Settings{
		StartTime: startTime,
		MachineID: func() (uint16, error) {
			return component.SonyflakeConfig.MachineId, nil
		},
	}
	snowflake := sonyflake.NewSonyflake(settings)
	if snowflake == nil {
		return errors.New("[components][core][sonyflake] new failed!")
	}
	component.Sonyflake = snowflake
	return nil
}
