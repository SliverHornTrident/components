//go:build tencent && cdn

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/tencent"
)

var (
	TencentCdn       map[string]*tencent.Cdn
	TencentCdnConfig config.TencentCdn
)
