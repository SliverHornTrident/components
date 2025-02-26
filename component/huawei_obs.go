//go:build huawei && obs

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	HuaweiObs       interfaces.Oss
	HuaweiObsConfig config.HuaweiObs
)
