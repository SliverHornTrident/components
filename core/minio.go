//go:build minio

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/minio"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Minio = new(_minio)

type _minio struct{}

func (c *_minio) Name() string {
	return "[components][core][minio]"
}

func (c *_minio) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Minio", &component.MinioConfig)
	if err != nil {
		return errors.Wrap(err, "viper Unmarshal failed")
	}
	return nil
}

func (c *_minio) IsPanic() bool {
	return false
}

func (c *_minio) ConfigName() string {
	return strings.Join([]string{"minio", gin.Mode(), "yaml"}, ".")
}

func (c *_minio) Initialization(ctx context.Context) error {
	client, err := minio.NewClient(component.MinioConfig)
	if err != nil {
		return errors.Wrap(err, "NewClient failed")
	}
	component.Minio = client
	return nil
}
