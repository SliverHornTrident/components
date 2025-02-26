//go:build aliyun

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Aliyun = new(_aliyun)

type _aliyun struct{}

func (c *_aliyun) Name() string {
	return "[components][core][aliyun]"
}

func (c *_aliyun) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Aliyun", &component.AliyunConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_aliyun) IsPanic() bool {
	return false
}

func (c *_aliyun) ConfigName() string {
	return strings.Join([]string{"aliyun", gin.Mode(), "yaml"}, ".")
}

func (c *_aliyun) Initialization(ctx context.Context) error {
	return nil
}
