package interfaces

import (
	"context"
)

type Oss interface {
	// Upload 通过 io.Reader 上传文件到oss
	Upload(ctx context.Context, name string, option *OssOption) (filepath string, filename string, err error)
	// DeleteFile 根据key删除文件
	DeleteFile(ctx context.Context, filename string) error
}

// OssConfig oss配置
type OssConfig interface {
	// Filename 根据文件名获取文件名
	Filename(name string) string
	// FileKey 根据文件名获取文件key
	FileKey(filename string) string
	// Filepath 根据文件key获取文件路径
	Filepath(key string) string
}
