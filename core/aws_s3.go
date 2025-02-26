//go:build aws && s3

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/aws"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var AwsS3 = new(awsS3)

type awsS3 struct{}

func (c *awsS3) Name() string {
	return "[components][core][aws]"
}

func (c *awsS3) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("AwsS3", &component.AwsS3Config)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.AwsS3Config.Aws = component.AwsConfig
	return nil
}

func (c *awsS3) IsPanic() bool {
	return false
}

func (c *awsS3) ConfigName() string {
	return strings.Join([]string{"asw", "s3", gin.Mode(), "yaml"}, ".")
}

func (c *awsS3) Initialization(ctx context.Context) error {
	s3, err := aws.NewS3(component.AwsS3Config)
	if err != nil {
		return err
	}
	component.AwsS3 = s3
	return nil
}
