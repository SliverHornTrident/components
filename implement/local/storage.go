//go:build local && storage

package local

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/pkg/errors"
	"io"
	"io/fs"
	"os"
)

type Storage struct {
	config config.LocalStorage
}

func NewStorage(config config.LocalStorage) (*Storage, error) {
	entity := &Storage{config: config}
	_, err := os.Stat(config.Path)
	if errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(config.Path, os.ModePerm)
		if err != nil {
			return nil, errors.Wrap(err, "[components][local][storage] create storage path failed!")
		}
	}
	return entity, nil
}

func (l *Storage) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	filename = l.config.Filename(name)
	filekey := l.config.FileKey(filename)
	filepath = l.config.Filepath(filekey)
	var out *os.File
	out, err = os.Create(filekey)
	if err != nil {
		return "", "", errors.Wrapf(err, "[components][local][storage] 创建文件失败!")
	}
	defer func() { // 创建文件流 defer 关闭
		_ = out.Close()
	}()
	_, err = io.Copy(out, option.GetReader()) // 传输(拷贝)文件
	if err != nil {
		return "", "", errors.Wrapf(err, "[components][local][storage] 传输(拷贝)文件失败!")
	}
	return filepath, filename, nil
}

func (l *Storage) DeleteFile(ctx context.Context, filename string) error {
	filekey := l.config.FileKey(filename)
	err := os.Remove(filekey)
	if err != nil {
		return errors.Wrap(err, "[components][local][storage] 删除文件失败!")
	}
	return nil
}
