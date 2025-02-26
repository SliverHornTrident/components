//go:build qiniu && kodo

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	QiniuKodo       interfaces.Oss
	QiniuKodoConfig config.QiniuKodo
)
