package tencent

import (
	"context"
	"github.com/SliverHornTrident/components/config"
	"testing"
)

func TestSes_SendMailWithTLS(t *testing.T) {
	type fields struct {
		Config config.TencentSes
	}
	type args struct {
		message *SesMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Config: config.TencentSes{
					Port:     465,
					Host:     "x",
					Username: "x",
					Password: "x",
					TLS:      true,
				},
			},
			args: args{
				message: &SesMessage{
					Body:    "<!DOCTYPE html>\\n<html>\\n<head>\\n<meta charset=\\\"utf-8\\\">\\n<title>hello world</title>\\n</head>\\n<body>\\n \" +\n\t\t\"<h1>我的第一个标题</h1>\\n    <p>我的第一个段落。</p>\\n</body>\\n</html>",
					Subject: "test",
					Tos:     []string{"503551462@qq.com"},
					Header:  nil,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ctx := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			s := NewSes(tt.fields.Config)
			if err := s.DialAndSend(ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("SendMailWithTLS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
