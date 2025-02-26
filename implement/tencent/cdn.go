//go:build tencent && cdn

package tencent

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/SliverHornTrident/components/config"
	"github.com/pkg/errors"
	"math/rand"
	"strconv"
	"time"
)

type Cdn struct {
	Config config.TencentCdnChildren
}

func NewCdn(config config.TencentCdnChildren) *Cdn {
	return &Cdn{Config: config}
}

// Signature 签名算法
// object 文件路径 不可以带 /
func (t *Cdn) Signature(object string) (url string, err error) {
	now := time.Now()
	key := t.Config.Key
	expire := now.Add(t.Config.Expire)
	timestamp := strconv.FormatInt(expire.Unix(), t.Config.ExpireBase)
	if key == "" {
		key = t.Config.BackupKey
	}
	switch t.Config.Mode {
	case "A":
		uid := "0"
		rand.New(rand.NewSource(time.Now().UnixNano()))
		length := rand.Intn(101)
		random := make([]rune, 0, length)
		runes := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
		for i := 0; i < length; i++ {
			rand.New(rand.NewSource(time.Now().UnixNano()))
			random[i] = runes[rand.Intn(len(runes))]
		}
		signature := fmt.Sprintf("%s-%s-%s-%s-%s", object, timestamp, string(random), uid, key)
		signature, err = t.Algorithm(signature)
		if err != nil {
			return "", errors.Wrap(err, "加密失败!")
		}
		return fmt.Sprintf("%s/%s?%s=%s-%s-%s-%s", t.Config.Domain, object, t.Config.SignatureKey, timestamp, string(random), uid, signature), nil
	case "B":
		signature := fmt.Sprintf("%s%s/%s", key, timestamp, object)
		signature, err = t.Algorithm(signature)
		if err != nil {
			return "", errors.Wrap(err, "加密失败!")
		}
		return fmt.Sprintf("%s/%s/%s/%s", t.Config.Domain, signature, timestamp, object), nil
	case "C":
		signature := fmt.Sprintf("%s/%s%s", key, object, timestamp)
		signature, err = t.Algorithm(signature)
		if err != nil {
			return "", errors.Wrap(err, "加密失败!")
		}
		return fmt.Sprintf("%s/%s/%s/%s", t.Config.Domain, signature, timestamp, object), nil
	case "D":
		signature := fmt.Sprintf("%s/%s%s", key, object, timestamp)
		signature, err = t.Algorithm(signature)
		if err != nil {
			return "", errors.Wrap(err, "加密失败!")
		}
		return fmt.Sprintf("%s/%s?%s=%s&%s=%s", t.Config.Domain, object, t.Config.SignatureKey, signature, t.Config.TimestampKey, timestamp), nil
	default:
		return "", errors.New("未知的鉴权模式!")
	}
}

func (t *Cdn) Algorithm(encryption string) (string, error) {
	switch t.Config.Algorithm {
	case "md5":
		hash := md5.New()
		hash.Write([]byte(encryption))
		return hex.EncodeToString(hash.Sum(nil)), nil
	case "sha256":
		hash := sha256.Sum256([]byte(encryption))
		return hex.EncodeToString(hash[:]), nil
	default:
		return "", errors.New("未知的鉴权算法!")
	}
}
