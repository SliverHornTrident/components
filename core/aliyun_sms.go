//go:build aliyun && sms

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

var AliyunSms = new(aliyunSms)

type aliyunSms struct{}

func (s *aliyunSms) Name() string {
	return "[components][core][aliyun][sms]"
}

func (s *aliyunSms) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("AliyunSms", &component.AliyunSmsConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	component.AliyunOssConfig.Aliyun = component.AliyunConfig
	return nil
}

func (s *aliyunSms) IsPanic() bool {
	return true
}

func (s *aliyunSms) ConfigName() string {
	return strings.Join([]string{"aliyun", "sms", gin.Mode(), "yaml"}, ".")
}

func (s *aliyunSms) Initialization(ctx context.Context) error {
	sms, err := aliyun.NewSms(component.AliyunSmsConfig)
	if err != nil {
		return err
	}
	component.AliyunSms = sms
	return nil
}
