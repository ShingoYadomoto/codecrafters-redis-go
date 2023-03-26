package resp

import (
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *command
		wantErr bool
	}{
		{
			args: args{
				b: []byte(`*1\r\n$4\r\nping\r\n`),
			},
			want: &command{
				cmd:  "PING",
				args: []string{},
			},
		},
		{
			args: args{
				b: []byte(`*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n`),
			},
			want: &command{
				cmd:  "ECHO",
				args: []string{"hey"},
			},
		},
		{
			args: args{
				b: []byte(`*2\r\n$4\r\nECHOOOOO\r\n$3\r\nhey\r\n`),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommand(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCommand() got = %v, want %v", got, tt.want)
			}
		})
	}
}
