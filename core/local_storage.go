//go:build local && storage

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/local"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var LocalStorage = new(localStorage)

type localStorage struct{}

func (c *localStorage) Name() string {
	return "[components][local][storage]"
}

func (c *localStorage) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("LocalStorage", &component.LocalStorageConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal config failed!")
	}
	return nil
}

func (c *localStorage) IsPanic() bool {
	return false
}

func (c *localStorage) ConfigName() string {
	return strings.Join([]string{"local", "storage", gin.Mode(), "yaml"}, ".")
}

func (c *localStorage) Initialization(ctx context.Context) error {
	storage, err := local.NewStorage(component.LocalStorageConfig)
	if err != nil {
		return err
	}
	component.LocalStorage = storage
	return nil
}
