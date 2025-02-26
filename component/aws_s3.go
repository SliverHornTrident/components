//go:build aws && s3

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	AwsS3       interfaces.Oss
	AwsS3Config config.AwsS3
)
