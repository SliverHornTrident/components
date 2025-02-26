//go:build jwt

package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

var SingleFlight = new(singleflight.Group)

type Jwt struct {
	SigningKey string
}

func NewJwt(key string) *Jwt {
	return &Jwt{SigningKey: key}
}

// Create 创建 token
func (j *Jwt) Create(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SigningKey))
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *Jwt) CreateTokenByOldToken(oldToken string, claims jwt.Claims) (string, error) {
	v, err, _ := SingleFlight.Do("Jwt:"+oldToken, func() (interface{}, error) {
		return j.Create(claims)
	})
	return v.(string), err
}

// Parse 解析 token
func (j *Jwt) Parse(tokenString string) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(j.SigningKey), nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token is invalid")
	}
	return nil, errors.New("token is nil")
}
