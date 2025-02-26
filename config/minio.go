//go:build minio

package config

import (
	"fmt"
	"path"
	"strings"
	"time"
)

type Minio struct {
	Path           string `json:"Path" yaml:"Path" mapstructure:"Path"`                               // 文件存储文件夹
	Prefix         string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`                         // 文件前缀 不填则为空
	Token          string `json:"Token" yaml:"Token" mapstructure:"Token"`                            // token
	Bucket         string `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`                         // 存储桶名称
	Domain         string `json:"Domain" yaml:"Domain" mapstructure:"Domain"`                         // 访问域名
	Endpoint       string `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`                   // 地域节点
	AccessKey      string `json:"AccessKey" yaml:"AccessKey" mapstructure:"AccessKey"`                // 秘钥AK
	SecretKey      string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"`                // 秘钥SK
	ExpirationTime string `json:"ExpirationTime" yaml:"ExpirationTime" mapstructure:"ExpirationTime"` // 过期时间
	UseSsl         bool   `json:"UseSsl" yaml:"UseSsl" mapstructure:"UseSsl"`                         // 是否使用ssl
}

func (c *Minio) Filename(name string) string {
	ext := path.Ext(name)
	if c.Prefix != "" {
		if strings.Contains(c.Prefix, "_") {
			return fmt.Sprintf("%s%d%s", c.Prefix, time.Now().UnixNano(), ext)
		}
		return fmt.Sprintf("%s_%d%s", c.Prefix, time.Now().UnixNano(), ext)
	}
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func (c *Minio) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *Minio) Filepath(key string) string {
	return path.Join(c.Domain, key)
}
