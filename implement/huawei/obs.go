//go:build huawei && obs

package huawei

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

type Obs struct {
	config config.HuaweiObs
	client *obs.ObsClient
}

func NewObs(config config.HuaweiObs) (interfaces.Oss, error) {
	entity := &Obs{config: config}
	client, err := obs.New(config.GetAccessKey(), config.GetAccessKey(), config.Endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "[components][huawei][obs] 创建client实例失败!")
	}
	entity.client = client
	return entity, nil
}

func (h *Obs) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	filename = h.config.Filename(name)
	filekey := h.config.FileKey(filename)
	filepath = h.config.Filepath(filename)
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: h.config.Bucket,
				Key:    filekey,
				Metadata: map[string]string{
					"filename": name,
				},
			},
		},
	}
	if option != nil {
		input.Body = option.GetReader()
		input.ContentType = option.GetHeader("content-type")
	}
	_, err = h.client.PutObject(input)
	if err != nil {
		return "", "", errors.Wrap(err, "[components][huawei][obs] 文件上传失败!")
	}
	return filepath, filename, nil
}

func (h *Obs) DeleteFile(ctx context.Context, filename string) error {
	filekey := h.config.FileKey(filename)
	_, err := h.client.DeleteObject(&obs.DeleteObjectInput{
		Bucket: h.config.Bucket,
		Key:    filekey,
	})
	if err != nil {
		return errors.Wrapf(err, "[components][huawei][obs][key:%s] 删除文件失败!", filekey)
	}
	return nil
}
