//go:build qiniu

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Qiniu = new(_qiniu)

type _qiniu struct{}

func (c *_qiniu) Name() string {
	return "[components][core][qiniu]"
}

func (c *_qiniu) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Qiniu", &component.QiniuConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_qiniu) IsPanic() bool {
	return false
}

func (c *_qiniu) ConfigName() string {
	return strings.Join([]string{"qiniu", gin.Mode(), "yaml"}, ".")
}

func (c *_qiniu) Initialization(ctx context.Context) error {
	return nil
}
