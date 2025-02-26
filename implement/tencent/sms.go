//go:build tencent && sms

package tencent

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type Sms struct {
	Client *sms.Client
	Config config.TencentSms
}

func NewSms(config config.TencentSms) (*Sms, error) {
	credential := config.NewCredential()
	clientProfile := profile.NewClientProfile()
	clientProfile.HttpProfile.ReqMethod = "POST"
	clientProfile.HttpProfile.Endpoint = config.Endpoint
	clientProfile.SignMethod = "HmacSHA1"
	client, err := sms.NewClient(credential, config.Region, clientProfile)
	if err != nil {
		return nil, errors.Wrap(err, "[components][tencent][sms] 链接失败!")
	}
	entity := &Sms{Config: config}
	entity.Client = client
	return entity, nil
}

func (t *Sms) SendSms(ctx context.Context, captcha int, phones ...*string) ([]*string, error) {
	request := sms.NewSendSmsRequest()
	request.PhoneNumberSet = phones
	request.SmsSdkAppId = &t.Config.AppId
	request.SignName = &t.Config.SignName
	request.TemplateId = &t.Config.TemplateId
	codes := t.Config.TemplateParam(captcha, phones...)
	request.TemplateParamSet = codes
	response, err := t.Client.SendSmsWithContext(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "[components][tencent][sms] 发送失败!")
	}
	for i := 0; i < len(response.Response.SendStatusSet); i++ {
		if response.Response.SendStatusSet[i].Code != nil {
			code := *response.Response.SendStatusSet[i].Code
			switch code {
			case "LimitExceeded.PhoneNumberThirtySecondLimit":
				return nil, errors.New("操作过于频繁,请30秒后重试!")
			case "LimitExceeded.PhoneNumberOneHourLimit":
				return nil, errors.New("操作过于频繁,请1小时后重试!")
			case "LimitExceeded.PhoneNumberDailyLimit":
				return nil, errors.New("操作过于频繁,请1天后重试!")
			case "FailedOperation.InsufficientBalanceInSmsPackage":
				return nil, errors.New("发送短信失败!")
			case "Ok":
				return codes, nil
			default:
				return nil, errors.New("未知错误!")
			}
		}
	}
	return nil, errors.New("未知错误!")
}
