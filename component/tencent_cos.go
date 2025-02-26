//go:build tencent && cos

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentCos       *tencent.Cos
	TencentCosConfig config.TencentCos
)
