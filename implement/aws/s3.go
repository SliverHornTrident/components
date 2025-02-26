//go:build aws && s3

package aws

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/interfaces"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

type S3 struct {
	config  config.AwsS3
	session *session.Session
}

func NewS3(config config.AwsS3) (*S3, error) {
	entity := &S3{config: config}
	_session, err := session.NewSession(&aws.Config{
		Region:           &config.Region,
		Endpoint:         &config.Endpoint, // minio在这里设置地址,可以兼容
		DisableSSL:       &config.DisableSsl,
		S3ForcePathStyle: &config.S3ForcePathStyle,
		Credentials: credentials.NewStaticCredentials(
			config.SecretID,
			config.SecretKey,
			"",
		),
	})
	if err != nil {
		return nil, errors.Wrap(err, "[components][aws][s3] 获取 session 失败!")
	}
	entity.session = _session
	return entity, nil
}

func (a *S3) Upload(ctx context.Context, name string, option *interfaces.OssOption) (filepath string, filename string, err error) {
	filename = a.config.Filename(name)
	filekey := a.config.FileKey(name)
	filepath = a.config.Filepath(name)
	uploader := s3manager.NewUploader(a.session)
	output, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.config.Bucket),
		Key:    aws.String(filekey),
		Body:   option.GetReader(),
	})
	if err != nil {
		return "", "", errors.Wrapf(err, "[oss][aws s3][output:%v]文件上传失败!", output)
	}
	return filename, filepath, nil
}

func (a *S3) DeleteFile(ctx context.Context, filename string) error {
	service := s3.New(a.session)
	filekey := a.config.FileKey(filename)
	_, err := service.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &a.config.Bucket,
		Key:    &filekey,
	})
	if err != nil {
		return errors.Wrap(err, "[components][aws][s3] 文件删除失败!")
	}
	return nil
}
