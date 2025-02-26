//go:build aliyun && sms

package aliyun

import (
	"github.com/SliverHornTrident/components/config"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/pkg/errors"
)

type Sms struct {
	client *dysmsapi.Client
	config config.AliyunSms
}

func NewSms(config config.AliyunSms) (*Sms, error) {
	entity := &Sms{
		config: config,
	}
	accessKeyId := config.GetAccessKeyId()
	accessKeySecret := config.GetAccessKeySecret()
	client, err := dysmsapi.NewClient(&openapi.Config{
		Endpoint:        &config.Endpoint,
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	})
	if err != nil {
		return nil, err
	}
	entity.client = client
	return entity, nil
}

func (a *Sms) SendSms(phone string) (code string, err error) {
	var param string
	param, code, err = a.config.TemplateParam(phone)
	if err != nil {
		return "", errors.Wrap(err, "[components][aliyun][sms]构造短信模板参数失败!")
	}
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &phone,
		SignName:      &a.config.SignName,
		TemplateCode:  &a.config.TemplateCode,
		TemplateParam: &param,
	}
	response, err := a.client.SendSms(request)
	if err != nil {
		return "", errors.Wrap(err, "[components][aliyun][sms]发送短信失败！")
	}
	if response != nil && response.Body != nil {
		if *response.Body.Code != "OK" {
			return "", errors.Errorf("[components][aliyun][sms]", *response.Body.Message)
		}
	}
	return code, nil
}
