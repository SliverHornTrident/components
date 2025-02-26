//go:build local && storage

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	LocalStorage       interfaces.Oss
	LocalStorageConfig config.LocalStorage
)
