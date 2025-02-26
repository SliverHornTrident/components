//go:build tencent && sms

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentSms       *tencent.Sms
	TencentSmsConfig config.TencentSms
)
