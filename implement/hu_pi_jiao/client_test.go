package hu_pi_jiao

import (
	"github.com/SliverHornTrident/components/config"
	"testing"
)

// BenchmarkClient_Sign
func BenchmarkClient_Sign(b *testing.B) {
	c := &Client{
		Config: config.HuPiJiaoChildren{
			AppId:       "1",
			Gateway:     "2",
			AppSecret:   "3",
			ReturnUrl:   "4",
			NotifyUrl:   "5",
			CallbackUrl: "6",
		},
	}
	params := map[string]string{
		"1": "1",
		"2": "2",
	}
	for i := 0; i < b.N; i++ {
		c.Sign(params)
	}
}

// BenchmarkClient_SignV2
func BenchmarkClient_SignV2(b *testing.B) {
	c := &Client{
		Config: config.HuPiJiaoChildren{
			AppId:       "1",
			Gateway:     "2",
			AppSecret:   "3",
			ReturnUrl:   "4",
			NotifyUrl:   "5",
			CallbackUrl: "6",
		},
	}
	params := map[string]string{
		"1": "1",
		"2": "2",
	}
	for i := 0; i < b.N; i++ {
		c.SignV2(params)
	}
}
