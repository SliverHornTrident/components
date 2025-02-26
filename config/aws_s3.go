//go:build aws && s3

package config

import (
	"fmt"
	"path"
	"strings"
	"time"
)

type AwsS3 struct {
	Path             string `json:"Path" yaml:"Path" mapstructure:"Path"`                                     // 文件存储文件夹
	Bucket           string `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`                               // 存储桶名称
	Domain           string `json:"Domain" yaml:"Domain" mapstructure:"Domain"`                               // 访问域名
	Region           string `json:"Region" yaml:"Region" mapstructure:"Region"`                               // 区域
	Prefix           string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`                               // 文件前缀 不填则为空
	Endpoint         string `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`                         // 地域节点
	SecretID         string `json:"SecretID" yaml:"SecretID" mapstructure:"SecretID"`                         // 密钥ID
	SecretKey        string `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"`                      // 密钥Key
	DisableSsl       bool   `json:"DisableSsl" yaml:"DisableSsl" mapstructure:"DisableSsl"`                   // 是否禁用ssl
	S3ForcePathStyle bool   `json:"S3ForcePathStyle" yaml:"S3ForcePathStyle" mapstructure:"S3ForcePathStyle"` // S3强制路径样式
	Aws              Aws    `json:"Aws" yaml:"Aws" mapstructure:"Aws"`                                        // AWS配置
}

func (c *AwsS3) Filename(name string) string {
	ext := path.Ext(name)
	if c.Prefix != "" {
		if strings.Contains(c.Prefix, "_") {
			return fmt.Sprintf("%s%d%s", c.Prefix, time.Now().UnixNano(), ext)
		}
		return fmt.Sprintf("%s_%d%s", c.Prefix, time.Now().UnixNano(), ext)
	}
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func (c *AwsS3) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *AwsS3) Filepath(key string) string {
	return path.Join(c.Domain, key)
}
