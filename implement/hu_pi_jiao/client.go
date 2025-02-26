//go:build hu_pi_jiao

package hu_pi_jiao

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/SliverHornTrident/components/config"
	"github.com/SliverHornTrident/components/implement/hu_pi_jiao/request"
	"github.com/SliverHornTrident/components/implement/hu_pi_jiao/response"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Config config.HuPiJiaoChildren
}

func NewClient(config config.HuPiJiaoChildren) *Client {
	return &Client{Config: config}
}

// Pay 支付
func (c *Client) Pay(ctx context.Context, info request.Pay) (*response.Pay, error) {
	now := time.Now().Unix()
	info.Time = now
	info.Appid = c.Config.AppId
	info.NonceStr = strconv.FormatInt(now, 10)
	info.ReturnUrl = c.Config.ReturnUrl
	info.NotifyUrl = c.Config.NotifyUrl
	info.CallbackUrl = c.Config.CallbackUrl
	values := make(url.Values, 17)

	var params map[string]string
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] json marshal failed!")
	}
	err = json.Unmarshal(bytes, &params)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] json unmarshal failed!")
	}
	for k, v := range params {
		values.Add(k, v)
	}
	values.Add("hash", c.SignV2(params))
	result, err := http.PostForm(c.Config.PayGateway, values)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = result.Body.Close()
	}()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] read response body failed!")
	}
	var data response.Pay
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] json unmarshal failed!")
	}
	return &data, err
}

// OrderQuery 订单查询
func (c *Client) OrderQuery(ctx context.Context, info request.OrderQuery) (*response.OrderQuery, error) {
	now := time.Now().Unix()
	info.Time = now
	info.AppId = c.Config.AppId
	info.NonceStr = strconv.FormatInt(now, 10)
	values := make(url.Values, 17)

	var params map[string]string
	bytes, err := json.Marshal(info)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] json marshal failed!")
	}
	err = json.Unmarshal(bytes, &params)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] json unmarshal failed!")
	}
	for k, v := range params {
		values.Add(k, v)
	}
	values.Add("hash", c.SignV2(params))
	result, err := http.PostForm(c.Config.OrderQueryGateway, values)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = result.Body.Close()
	}()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, errors.Wrap(err, "[components][hu_pi_jiao] read response body failed!")
	}
	var data response.OrderQuery
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "[components][hu_pi_jiao][data:%s] json unmarshal failed!", string(body))
	}
	return &data, err
}

// Sign 签名方法
func (c *Client) Sign(params map[string]string) string {
	var data string
	keys := make([]string, 0)
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		data = fmt.Sprintf("%s%s=%s&", data, k, params[k])
	}
	data = strings.Trim(data, "&")
	data = fmt.Sprintf("%s%s", data, c.Config.AppSecret)
	m := md5.New()
	m.Write([]byte(data))
	sign := fmt.Sprintf("%x", m.Sum(nil))
	return sign
}

// SignV2 签名方法 v2
func (c *Client) SignV2(params map[string]string) string {
	var data string
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	length := len(keys)
	for i := 0; i < length; i++ {
		data += keys[i] + "=" + params[keys[i]] + "&"
	}
	data = strings.Trim(data, "&")
	data = fmt.Sprintf("%s%s", data, c.Config.AppSecret)
	m := md5.New()
	m.Write([]byte(data))
	sign := fmt.Sprintf("%x", m.Sum(nil))
	return sign
}
