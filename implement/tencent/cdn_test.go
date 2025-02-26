//go:build tencent && cdn

package tencent

import (
	"github.com/SliverHornTrident/components/config"
	"testing"
)

func TestCdn_Algorithm(t1 *testing.T) {
	type fields struct {
		Config config.TencentCdnChildren
	}
	type args struct {
		encryption string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "md5",
			fields: fields{
				Config: config.TencentCdnChildren{
					Algorithm: "md5",
				},
			},
			args: args{
				encryption: "dimtm5evg50ijsx2hvuwyfoiu65/test.jpg5e577978",
			},
			want: "7913fc0c5c9e92dd3633b7895152bbb2",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Cdn{
				Config: tt.fields.Config,
			}
			got, err := t.Algorithm(tt.args.encryption)
			if (err != nil) != tt.wantErr {
				t1.Errorf("Algorithm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("Algorithm() got = %v, want %v", got, tt.want)
			}
		})
	}
}
