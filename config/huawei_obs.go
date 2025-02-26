//go:build huawei && obs

package config

import (
	"fmt"
	"net/url"
	"path"
	"time"
)

type HuaweiObs struct {
	Path      string `json:"Path" yaml:"Path" mapstructure:"Path"`                // 文件存储文件夹
	Prefix    string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`          // 文件前缀 不填则为空
	Bucket    string `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`          // 存储桶名称
	Domain    string `json:"Domain" yaml:"Domain" mapstructure:"Domain"`          // 访问域名
	Endpoint  string `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`    // 地域节点
	AccessKey string `json:"AccessKey" yaml:"AccessKey" mapstructure:"AccessKey"` // 访问密钥
	SecretKey string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 密钥
	Huawei    Huawei `json:"Huawei" yaml:"Huawei" mapstructure:"Huawei"`
}

func (c *HuaweiObs) Filename(name string) string {
	if c.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), name)
	}
	return fmt.Sprintf("%s%d_%s", c.Prefix, time.Now().Local().Unix(), name)
}

func (c *HuaweiObs) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *HuaweiObs) Filepath(key string) string {
	link, _ := url.JoinPath(c.Domain, key)
	return link
}

func (c *HuaweiObs) GetAccessKey() string {
	if c.AccessKey == "" {
		return c.Huawei.AccessKey
	}
	return c.AccessKey
}

func (c *HuaweiObs) GetSecretKey() string {
	if c.SecretKey == "" {
		return c.Huawei.SecretKey
	}
	return c.SecretKey
}
