//go:build tencent && cos

package tencent

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/pkg/errors"
	"net/http"
)

type Cos struct {
	Config config.TencentCos
	Client *cos.Client
}

func NewCos(config config.TencentCos) (*Cos, error) {
	entity := &Cos{Config: config}
	uri, err := config.Uri()
	if err != nil {
		return nil, err
	}
	entity.Client = cos.NewClient(uri, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretId,
			SecretKey: config.SecretKey,
		},
	})
	return entity, nil
}

func (c *Cos) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filekey string, err error) {
	filename := c.Config.Filename(name)
	filekey = c.Config.FileKey(filename)
	filepath = c.Config.Filepath(filekey)
	options := &cos.ObjectPutOptions{}
	_, err = c.Client.Object.Put(ctx, filekey, option.GetReader(), options)
	if err != nil {
		return "", "", errors.Wrap(err, "[components][tencent][cos] 上传文件失败!")
	}
	return filepath, filekey, nil
}

func (c *Cos) DeleteFile(ctx context.Context, filename string) error {
	filekey := c.Config.FileKey(filename)
	_, err := c.Client.Object.Delete(ctx, filekey)
	if err != nil {
		return errors.Wrapf(err, "[components][tencent][cos][key:%s] 删除文件失败!", filekey)
	}
	return nil
}
