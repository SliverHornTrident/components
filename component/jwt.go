//go:build jwt

package component

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/jwt"
)

var (
	Jwt       *jwt.Jwt
	JwtConfig config.Jwt
)
