//go:build qiniu && kodo

package qiniu

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/pkg/errors"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type Kodo struct {
	config        config.QiniuKodo
	credentials   *auth.Credentials
	storageConfig *storage.Config
}

func NewKodo(config config.QiniuKodo) (*Kodo, error) {
	accessKey := config.GetAccessKey()
	entity := &Kodo{config: config}
	entity.credentials = auth.New(accessKey, config.GetSecretKey())
	storageConfig := &storage.Config{
		UseHTTPS:      config.UseHttps,
		UseCdnDomains: config.UseCdnDomains,
	}
	region, err := storage.GetRegion(accessKey, config.Bucket) // 用来根据ak和bucket来获取空间相关的机房信息
	if err != nil {
		return nil, errors.Wrap(err, "[components][qiniu][kodo] 根据AccessKey和Bucket来获取空间相关的机房信息失败!")
	}
	storageConfig.Region = region
	return entity, nil
}

func (q *Kodo) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	putPolicy := storage.PutPolicy{Scope: q.config.Bucket}
	token := putPolicy.UploadToken(q.credentials)
	uploader := storage.NewFormUploader(q.storageConfig)
	filename = q.config.Filename(name)
	filekey := q.config.FileKey(filename)
	filepath = q.config.Filepath(filekey)
	info := option.GetFileInfo()
	var ret storage.PutRet
	err = uploader.Put(ctx, &ret, token, filekey, option.GetReader(), info.Size(), &storage.PutExtra{Params: map[string]string{"name": filename}})
	if err != nil {
		return "", "", errors.Wrap(err, "[components][qiniu][kodo] 上传文件失败!")
	}
	return filename, filepath, nil
}

func (q *Kodo) DeleteFile(ctx context.Context, filename string) error {
	filekey := q.config.FileKey(filename)
	err := storage.NewBucketManager(q.credentials, q.storageConfig).Delete(q.config.Bucket, filekey)
	if err != nil {
		return errors.Wrap(err, "[components][qiniu][kodo] 删除文件失败!")
	}
	return nil
}
