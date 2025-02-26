//go:build tencent && vod

package tencent

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/SliverHornTrident/components/config"
	"math/big"
	"time"
)

type Vod struct {
	Config config.TencentVod
	Client *vod.VodUploadClient
}

func NewVod(config config.TencentVod) *Vod {
	client := &vod.VodUploadClient{
		SecretId:  config.SecretId,
		SecretKey: config.SecretKey,
	}
	return &Vod{Config: config, Client: client}
}

func (v *Vod) Signature(expire time.Duration) (currentTimestamp int64, expireTimestamp int64, signature string) {
	var random string
	for i := 0; i < 6; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(10))
		random += result.String()
	}
	now := time.Now()
	currentTimestamp = now.Unix()
	expireTimestamp = now.Add(expire).Unix()
	original := fmt.Sprintf("secretId=%s&currentTimeStamp=%d&expireTime=%d&random=%s", v.Config.SecretId, currentTimestamp, expireTimestamp, random)
	sign := generateHmacSHA1(v.Config.SecretKey, original)
	sign = append(sign, []byte(original)...)
	signature = base64.StdEncoding.EncodeToString(sign)
	return currentTimestamp, expireTimestamp, signature
}
func generateHmacSHA1(secretToken, payloadBody string) []byte {
	mac := hmac.New(sha1.New, []byte(secretToken))
	sha1.New()
	mac.Write([]byte(payloadBody))
	return mac.Sum(nil)
}
