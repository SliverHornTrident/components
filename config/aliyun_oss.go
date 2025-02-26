//go:build aliyun && oss

package config

import (
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
)

type AliyunOss struct {
	Path            string `json:"Path" yaml:"Path" mapstructure:"Path"`                                  // 文件存储文件夹
	Prefix          string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`                            // 文件存储前缀
	Bucket          string `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`                            // 存储桶名称
	Domain          string `json:"Domain" yaml:"Domain" mapstructure:"Domain"`                            // 访问域名
	Endpoint        string `json:"Endpoint" yaml:"Endpoint" mapstructure:"Endpoint"`                      // 地域节点
	AccessKeyId     string `json:"AccessKeyId" yaml:"AccessKeyId" mapstructure:"AccessKeyId"`             // 阿里云对象存储 访问密钥Id
	AccessKeySecret string `json:"AccessKeySecret" yaml:"AccessKeySecret" mapstructure:"AccessKeySecret"` // 阿里云对象存储 访问密钥Secret
	Aliyun          Aliyun `json:"Aliyun" yaml:"-" mapstructure:"-"`                                      // 阿里云公共配置
}

func (c *AliyunOss) GetAccessKeyId() string {
	if c.AccessKeyId == "" {
		c.AccessKeyId = c.Aliyun.AccessKeyId
	}
	return c.AccessKeyId
}

func (c *AliyunOss) GetAccessKeySecret() string {
	if c.AccessKeySecret == "" {
		c.AccessKeySecret = c.Aliyun.AccessKeySecret
	}
	return c.AccessKeySecret
}

func (c *AliyunOss) Filename(name string) string {
	ext := path.Ext(name)
	if c.Prefix != "" {
		if strings.Contains(c.Prefix, "_") {
			return fmt.Sprintf("%s%d%s", c.Prefix, time.Now().UnixNano(), ext)
		}
		return fmt.Sprintf("%s_%d%s", c.Prefix, time.Now().UnixNano(), ext)
	}
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func (c *AliyunOss) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *AliyunOss) Filepath(key string) string {
	link, _ := url.JoinPath(c.Domain, key)
	return link
}
