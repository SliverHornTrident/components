//go:build qiniu && kodo

package config

import (
	"fmt"
	"path"
	"strings"
	"time"
)

type QiniuKodo struct {
	Path          string `json:"Path" yaml:"Path" mapstructure:"Path"`                            // 文件存储文件夹
	Prefix        string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`                      // 文件前缀 不填则为空
	Bucket        string `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`                      // 空间名称
	Domain        string `json:"Domain" yaml:"Domain" mapstructure:"Domain"`                      // CDN加速域名
	AccessKey     string `json:"AccessKey" yaml:"AccessKey" mapstructure:"AccessKey"`             // 秘钥AK
	SecretKey     string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"`             // 秘钥SK
	UseHttps      bool   `json:"UseHttps" yaml:"UseHttps" mapstructure:"UseHttps"`                // 是否使用https
	UseCdnDomains bool   `json:"UseCdnDomains" yaml:"UseCdnDomains" mapstructure:"UseCdnDomains"` // 是否使用cdn加速域名
	Qiniu         Qiniu  `json:"Qiniu" yaml:"Qiniu" mapstructure:"Qiniu"`                         // 七牛云配置
}

func (c *QiniuKodo) GetAccessKey() string {
	if c.AccessKey == "" {
		return c.Qiniu.AccessKey
	}
	return c.AccessKey
}

func (c *QiniuKodo) GetSecretKey() string {
	if c.SecretKey == "" {
		return c.Qiniu.SecretKey
	}
	return c.SecretKey
}

func (c *QiniuKodo) Filename(name string) string {
	ext := path.Ext(name)
	if c.Prefix != "" {
		if strings.Contains(c.Prefix, "_") {
			return fmt.Sprintf("%s%d%s", c.Prefix, time.Now().UnixNano(), ext)
		}
		return fmt.Sprintf("%s_%d%s", c.Prefix, time.Now().UnixNano(), ext)
	}
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func (c *QiniuKodo) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *QiniuKodo) Filepath(key string) string {
	return path.Join(c.Domain, key)
}
