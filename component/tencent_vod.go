//go:build tencent && vod

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentVod       *tencent.Vod
	TencentVodConfig config.TencentVod
)
