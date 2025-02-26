//go:build cron

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/cron"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

var Cron = new(_cron)

type _cron struct{}

func (c *_cron) Name() string {
	return "[components][core][cron]"
}

func (c *_cron) Viper(viper *viper.Viper) error {
	return nil
}

func (c *_cron) IsPanic() bool {
	return false
}

func (c *_cron) ConfigName() string {
	return strings.Join([]string{"cron", gin.Mode(), "yaml"}, ".")
}

func (c *_cron) Initialization(ctx context.Context) error {
	component.Cron = cron.NewCron()
	return nil
}
