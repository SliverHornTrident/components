//go:build aliyun && sms

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/aliyun"
)

var (
	AliyunSms       *aliyun.Sms
	AliyunSmsConfig config.AliyunSms
)
