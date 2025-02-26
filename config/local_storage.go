//go:build local && storage

package config

import (
	"fmt"
	"net/url"
	"path"
	"time"
)

type LocalStorage struct {
	Path   string `json:"Path" yaml:"Path" mapstructure:"Path"`       // 存储路径
	Prefix string `json:"Prefix" yaml:"Prefix" mapstructure:"Prefix"` // 文件前缀 不填则为空
	Domain string `json:"Domain" yaml:"Domain" mapstructure:"Domain"` // 访问路径
}

func (c *LocalStorage) Filename(name string) string {
	if c.Prefix == "" {
		return fmt.Sprintf("%d_%s", time.Now().Local().Unix(), name)
	}
	return fmt.Sprintf("%s%d_%s", c.Prefix, time.Now().Local().Unix(), name)
}

func (c *LocalStorage) FileKey(filename string) string {
	return path.Join(c.Path, filename)
}

func (c *LocalStorage) Filepath(key string) string {
	link, _ := url.JoinPath(c.Domain, key)
	return link
}
