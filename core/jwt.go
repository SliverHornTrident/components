//go:build jwt

package core

import (
	"context"
	"github.com/SliverHornTrident/components/component"
	"github.com/SliverHornTrident/components/implement/jwt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

var Jwt = new(_jwt)

type _jwt struct{}

func (c *_jwt) Name() string {
	return "[components][core][jwt]"
}

func (c *_jwt) Viper(viper *viper.Viper) error {
	err := viper.UnmarshalKey("Jwt", &component.JwtConfig)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (c *_jwt) IsPanic() bool {
	return false
}

func (c *_jwt) ConfigName() string {
	return strings.Join([]string{"jwt", gin.Mode(), "yaml"}, ".")
}

func (c *_jwt) Initialization(ctx context.Context) error {
	component.Jwt = jwt.NewJwt(component.JwtConfig.SigningKey)
	return nil
}
