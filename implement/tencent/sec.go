//go:build tencent && sec

package tencent

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/SliverHornTrident/components/config"
	"github.com/pkg/errors"
	captcha "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/captcha/v20190722"
	"time"
)

type Sec struct {
	Config config.TencentSec
	Client *captcha.Client
}

func NewSec(config config.TencentSec) (*Sec, error) {
	credential := config.NewCredential()
	clientProfile := config.ClientProfile()
	client, err := captcha.NewClient(credential, config.Region, clientProfile)
	if err != nil {
		return nil, errors.Wrap(err, "链接失败!")
	}
	return &Sec{Config: config, Client: client}, nil
}

func (s *Sec) EncryptAppId(expire time.Duration) (id string, err error) {
	appSecretKey := []byte(s.Config.AppSecretKey)
	remainder := 32 % len(appSecretKey)
	key := append(appSecretKey, appSecretKey[:remainder]...)
	currentTimeStamp := time.Now().Unix()
	expireTime := int64(expire.Seconds())
	text := fmt.Sprintf("%d&%d&%d", s.Config.CaptchaAppId, currentTimeStamp, expireTime)
	iv := make([]byte, 16)
	_, err = rand.Read(iv)
	if err != nil {
		return "", errors.Wrap(err, "生成iv失败!")
	}
	id, err = AesCBCEncrypt([]byte(text), key, iv)
	if err != nil {
		return "", errors.Wrap(err, "加密失败!")
	}
	return id, nil
}

// AesCBCEncrypt AES/CBC/PKCS7Padding 加密
func AesCBCEncrypt(text []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key) // AES
	if err != nil {
		return "", errors.New("key error")
	}
	text = PaddingPKCS7(text, aes.BlockSize) // PKCS7 填充

	// CBC 加密
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(text, text)
	return base64.StdEncoding.EncodeToString(append(iv, text...)), nil
}

// AesCBCDecrypt AES/CBC/PKCS7Padding 解密(unknown test)
func AesCBCDecrypt(text []byte, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key) // AES
	if err != nil {
		return "", err
	}
	if len(text)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}
	{
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(text, text)
	} // CBC 解密
	empty := UnPaddingPKCS7(text) // PKCS7 反填充
	return base64.StdEncoding.EncodeToString(append(iv, empty...)), nil
}

// UnPaddingPKCS7 反填充
func UnPaddingPKCS7(text []byte) []byte {
	length := len(text)
	if length == 0 {
		return text
	}
	unPadding := int(text[length-1])
	return text[:(length - unPadding)]
}

// PaddingPKCS7 填充
func PaddingPKCS7(text []byte, size int) []byte {
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(text, paddingText...)
}
