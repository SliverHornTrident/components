//go:build jwt

package jwt

import (
	"github.com/SliverHornTrident/components/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Claims struct {
	UserId   uint64           `json:"user_id,omitempty"`
	RoleId   uint64           `json:"role_id,omitempty"`
	BufferAt *jwt.NumericDate `json:"buf,omitempty"`
	jwt.RegisteredClaims
}

func (c *Claims) GetUserId() uint64 {
	if c.UserId == 0 {
		id, _ := strconv.ParseUint(c.ID, 10, 64)
		return id
	}
	return c.UserId
}

func (c *Claims) GetRoleId() uint64 {
	return c.RoleId
}

type ClaimsOption func(claims *Claims)

func NewClaims(config config.Jwt, options ...ClaimsOption) *Claims {
	now := time.Now()
	expiresAt := now.Add(config.ExpiresAt)
	bufferAt := expiresAt.Add(-(config.ExpiresAt - config.BufferAt))
	claims := &Claims{
		BufferAt: jwt.NewNumericDate(bufferAt),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(now.Add(time.Duration(-1000))),
		},
	}
	for i := 0; i < len(options); i++ {
		options[i](claims)
	}
	return claims
}

func ClaimsWithUserId(userId uint64) ClaimsOption {
	return func(claims *Claims) {
		claims.UserId = userId
	}
}

func ClaimsWithRoleId(roleId uint64) ClaimsOption {
	return func(claims *Claims) {
		claims.RoleId = roleId
	}
}

// NewClaimsByGin 从gin.Context中获取Claims
// you need use jwt middleware before this function and gin.Context.Set("claims", claims) in middleware
func NewClaimsByGin(ctx *gin.Context) (claims *Claims, err error) {
	value, exists := ctx.Get("claims")
	if !exists {
		return nil, errors.New("未启用jwt中间件!")
	}
	claims, exists = value.(*Claims)
	if !exists {
		return nil, errors.New("类型断言失败!")
	}
	return claims, nil
}
