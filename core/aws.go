//go:build aws

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Aws = new(_aws)

type _aws struct{}

func (c *_aws) Name() string {
	return "[components][core][aws]"
}

func (c *_aws) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Aws", &component.AwsConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	return nil
}

func (c *_aws) IsPanic() bool {
	return false
}

func (c *_aws) ConfigName() string {
	return strings.Join([]string{"asw", gin.Mode(), "yaml"}, ".")
}

func (c *_aws) Initialization(ctx context.Context) error {
	return nil
}
