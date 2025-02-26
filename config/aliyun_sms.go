//go:build aliyun && sms

package config

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"strconv"
)

type AliyunSms struct {
	Aliyun          Aliyun `json:"Aliyun" yaml:"-" mapstructure:"-"`
	SignName        string `json:"SignName" yaml:"SignName" mapstructure:"SignName"`                      // 签名
	Endpoint        string `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`                      // 地域节点
	TemplateCode    string `json:"TemplateCode" yaml:"TemplateCode" mapstructure:"TemplateCode"`          // 模板编号
	AccessKeyId     string `json:"AccessKeyId" yaml:"AccessKeyId" mapstructure:"AccessKeyId"`             // 阿里云短信 访问密钥Id
	AccessKeySecret string `json:"AccessKeySecret" yaml:"AccessKeySecret" mapstructure:"AccessKeySecret"` // 阿里云短信 访问密钥Secret
}

func (c *AliyunSms) GetAccessKeyId() string {
	if c.AccessKeyId == "" {
		c.AccessKeyId = c.Aliyun.AccessKeyId
	}
	return c.AccessKeyId
}

func (c *AliyunSms) GetAccessKeySecret() string {
	if c.AccessKeySecret == "" {
		c.AccessKeySecret = c.Aliyun.AccessKeySecret
	}
	return c.AccessKeySecret
}

func (c *AliyunSms) TemplateParam(phone string) (param, code string, err error) {
	type Response struct {
		Code string `json:"code"`
	}
	var r Response
	for i := 0; i < 6; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(10))
		r.Code += strconv.Itoa(int(result.Int64()))
	}
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", "", err
	}
	return string(bytes), r.Code, nil
}
