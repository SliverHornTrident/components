//go:build minio

package minio

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

type Client struct {
	config config.Minio
	client *minio.Client
}

func NewClient(config config.Minio) (*Client, error) {
	entity := &Client{config: config}
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKey, config.SecretKey, config.Token),
		Secure: config.UseSsl,
	})
	if err != nil {
		return nil, errors.Wrap(err, "[components][minio] 获取minio.Client对象失败!")
	}
	var exists bool
	ctx := context.Background()
	exists, err = client.BucketExists(ctx, config.Bucket)
	if err != nil {
		return nil, errors.Wrapf(err, "[components][minio][bucket:%s] 存储桶检查失败!", config.Bucket)
	}
	if !exists {
		var options minio.MakeBucketOptions
		err = client.MakeBucket(ctx, config.Bucket, options)
		if err != nil {
			return nil, errors.Wrapf(err, "[components][minio][bucket:%s] 存储桶创建失败!", config.Bucket)
		}
	}
	entity.client = client
	return entity, nil
}

func (c *Client) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	filename = c.config.Filename(name)
	filekey := c.config.FileKey(filename)
	filepath = c.config.Filepath(filekey)
	options := minio.PutObjectOptions{ContentType: option.GetHeader("content-type")}
	info := option.GetFileInfo()
	_, err = c.client.PutObject(ctx, c.config.Bucket, filekey, option.GetReader(), info.Size(), options)
	if err != nil {
		return "", "", errors.Wrap(err, "[components][minio] 文件上传失败!")
	}
	return filepath, filename, nil
}

func (c *Client) DeleteFile(ctx context.Context, filename string) error {
	options := minio.RemoveObjectOptions{GovernanceBypass: true}
	filekey := c.config.FileKey(filename)
	err := c.client.RemoveObject(ctx, c.config.Bucket, filekey, options)
	if err != nil {
		return errors.Wrap(err, "[components][minio] 删除文件失败!")
	}
	return nil
}
