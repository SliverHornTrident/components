//go:build qiniu && kodo

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/qiniu"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var QiniuKodo = new(qiniuKodo)

type qiniuKodo struct{}

func (c *qiniuKodo) Name() string {
	return "[components][core][qiniu][kodo]"
}

func (c *qiniuKodo) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("QiniuKodo", &component.QiniuKodoConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.QiniuKodoConfig.Qiniu = component.QiniuConfig
	return nil
}

func (c *qiniuKodo) IsPanic() bool {
	return false
}

func (c *qiniuKodo) ConfigName() string {
	return strings.Join([]string{"qiniu", "kodo", gin.Mode(), "yaml"}, ".")
}

func (c *qiniuKodo) Initialization(ctx context.Context) error {
	kodo, err := qiniu.NewKodo(component.QiniuKodoConfig)
	if err != nil {
		return err
	}
	component.QiniuKodo = kodo
	return nil
}
