//go:build hu_pi_jiao

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/hu_pi_jiao"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var HuPiJiao = new(huPiJiao)

type huPiJiao struct{}

func (h *huPiJiao) Name() string {
	return "[components][core][hu_pi_jiao]"
}

func (h *huPiJiao) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("HuPiJiao", &component.HuPiJiaoConfig)
	if err != nil {
		return errors.Wrap(err, "[components][core][hu_pi_jiao] viper unmarshal failed!")
	}
	return nil
}

func (h *huPiJiao) IsPanic() bool {
	return false
}

func (h *huPiJiao) ConfigName() string {
	return strings.Join([]string{"hu_pi_jiao", gin.Mode(), "yaml"}, ".")
}

func (h *huPiJiao) Initialization(ctx context.Context) error {
	length := len(component.HuPiJiaoConfig)
	component.HuPiJiao = make(map[string]*hu_pi_jiao.Client, length)
	for i := 0; i < length; i++ {
		client := hu_pi_jiao.NewClient(component.HuPiJiaoConfig[i])
		component.HuPiJiao[component.HuPiJiaoConfig[i].Name] = client
	}
	return nil
}
