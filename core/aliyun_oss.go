//go:build aliyun && oss

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/aliyun"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var AliyunOss = new(aliyunOss)

type aliyunOss struct{}

func (c *aliyunOss) Name() string {
	return "[components][core][aliyun][oss]"
}

func (c *aliyunOss) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("AliyunOss", &component.AliyunOssConfig)
	if err != nil {
		return errors.Wrap(err, "viper unmarshal failed!")
	}
	component.AliyunOssConfig.Aliyun = component.AliyunConfig
	return nil
}

func (c *aliyunOss) IsPanic() bool {
	return false
}

func (c *aliyunOss) ConfigName() string {
	return strings.Join([]string{"aliyun", "oss", gin.Mode(), "yaml"}, ".")
}

func (c *aliyunOss) Initialization(ctx context.Context) error {
	oss, err := aliyun.NewOss(component.AliyunOssConfig)
	if err != nil {
		return err
	}
	component.AliyunOss = oss
	return nil
}
