//go:build tencent && sec

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentSec       *tencent.Sec
	TencentSecConfig config.TencentSec
)
