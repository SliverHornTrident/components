//go:build aliyun && oss

package aliyun

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"
)

type Oss struct {
	config config.AliyunOss
	client *oss.Client
	bucket *oss.Bucket
}

func NewOss(config config.AliyunOss) (interfaces.Oss, error) {
	entity := &Oss{config: config}
	client, err := oss.New(config.Endpoint, config.GetAccessKeyId(), config.GetAccessKeySecret()) // 创建 oss.Client 实例。
	if err != nil {
		return nil, errors.Wrap(err, "[components][aliyun][oss] 创建client实例失败!")
	}
	entity.client = client
	entity.bucket, err = entity.client.Bucket(config.Bucket) // 获取存储空间
	if err != nil {
		return nil, errors.Wrap(err, "[components][aliyun][oss] 获取存储空间失败!")
	}
	return entity, nil
}

func (a *Oss) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	filename = a.config.Filename(name)
	filekey := a.config.FileKey(filename)
	filepath = a.config.Filepath(filekey)
	err = a.bucket.PutObject(filekey, option.GetReader(), oss.Meta("filename", filename))
	if err != nil {
		return "", "", errors.Wrap(err, "[components][aliyun][oss] 上传文件失败!")
	}
	return filepath, filename, nil
}

func (a *Oss) DeleteFile(ctx context.Context, filename string) error {
	filekey := a.config.FileKey(filename)
	err := a.bucket.DeleteObject(filekey)
	if err != nil {
		return errors.Wrapf(err, "[components][aliyun][oss][key:%s]删除文件失败!", filekey)
	}
	return nil
}
