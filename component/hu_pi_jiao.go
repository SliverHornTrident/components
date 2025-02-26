//go:build hu_pi_jiao

package component

import (
	"github.com/SliverHornTrident/components/config"
	implement "github.com/SliverHornTrident/components/implement/hu_pi_jiao"
)

var (
	HuPiJiao       map[string]*implement.Client
	HuPiJiaoConfig config.HuPiJiao
)
