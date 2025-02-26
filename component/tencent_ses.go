//go:build tencent && ses

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentSes       *tencent.Ses
	TencentSesConfig config.TencentSes
)
