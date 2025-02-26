package tencent

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	AppSecretKey := []byte("1234567891011121314151516")
	remainder := 32 % len(AppSecretKey)
	key := append(AppSecretKey, AppSecretKey[:remainder]...)

	CaptchaAppId := "123456789"
	curTime := 1710144972
	expireTime := 86400

	plaintext := []byte(fmt.Sprintf("%s&%d&%d", CaptchaAppId, curTime, expireTime))

	iv := []byte("0123456789012345")

	ciphertext, err := AesCBCEncrypt(plaintext, key, iv)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Ciphertext (Base64):", ciphertext)
}
