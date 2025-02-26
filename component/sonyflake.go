//go:build sonyflake

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/sony/sonyflake"
)

var (
	Sonyflake       *sonyflake.Sonyflake
	SonyflakeConfig config.Sonyflake
)
