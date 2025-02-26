//go:build minio

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	Minio       interfaces.Oss
	MinioConfig config.Minio
)
