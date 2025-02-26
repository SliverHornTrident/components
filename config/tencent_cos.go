//go:build tencent && cos

package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/url"
	"path"
	"strings"
	"time"
)

type TencentCos struct {
	Path      string  `json:"Path" yaml:"Path" mapstructure:"Path"`                // 文件存储文件夹
	Prefix    string  `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"`          // 文件名前缀
	Domain    string  `json:"Domain" yaml:"Domain" mapstructure:"Domain"`          // 访问域名
	Private   string  `json:"Private" yaml:"Private" mapstructure:"Private"`       // 私有域名
	Bucket    string  `json:"Bucket" yaml:"Bucket" mapstructure:"Bucket"`          // 存储桶
	SecretId  string  `json:"SecretId" yaml:"SecretId" mapstructure:"SecretId"`    // 访问密钥 Id
	SecretKey string  `json:"SecretKey" yaml:"SecretKey" mapstructure:"SecretKey"` // 访问密钥 Secret
	Tencent   Tencent `json:"Tencent" yaml:"-" mapstructure:"-"`
}

func (c *TencentCos) NewCredential() *common.Credential {
	if c.SecretId != "" {
		c.Tencent.SecretId = c.SecretId
	}
	if c.SecretKey != "" {
		c.Tencent.SecretKey = c.SecretKey
	}
	return c.Tencent.NewCredential()
}

func (c *TencentCos) Filename(name string) string {
	ext := path.Ext(name)
	if c.Prefix != "" {
		if strings.Contains(c.Prefix, "_") {
			return fmt.Sprintf("%s%d%s", c.Prefix, time.Now().UnixNano(), ext)
		}
		return fmt.Sprintf("%s_%d%s", c.Prefix, time.Now().UnixNano(), ext)
	}
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

func (c *TencentCos) FileKey(filename string) string {
	if strings.Contains(filename, c.Path) {
		return filename
	}
	return path.Join(c.Path, filename)
}

func (c *TencentCos) Filepath(key string) string {
	link, _ := url.JoinPath(c.Domain, key)
	return link
}

func (c *TencentCos) Uri() (*cos.BaseURL, error) {
	domain, err := url.Parse(c.Domain)
	if err != nil {
		return nil, errors.Wrapf(err, "[components][tencent][cos] 解析域名[%s]失败!", c.Domain)
	}
	return &cos.BaseURL{BucketURL: domain}, nil
}
