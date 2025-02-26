//go:build aliyun && oss

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
)

var (
	AliyunOss       interfaces.Oss
	AliyunOssConfig config.AliyunOss
)
