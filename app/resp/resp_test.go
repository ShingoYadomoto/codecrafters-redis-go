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
				b: []byte("*1\r\n$4\r\nping\r\n"),
			},
			want: &command{
				cmd:     "PING",
				argsStr: "",
				argsLen: 0,
			},
		},
		{
			args: args{
				b: []byte("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"),
			},
			want: &command{
				cmd:     "ECHO",
				argsStr: "$3\r\nhey\r\n",
				argsLen: 1,
			},
		},
		{
			args: args{
				b: []byte("*2\r\n$4\r\nECHOOOOO\r\n$3\r\nhey\r\n"),
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

func Test_command_Response(t *testing.T) {
	type fields struct {
		cmd     string
		argsStr string
		argsLen int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "ping -> pong",
			fields: fields{
				cmd:     "PING",
				argsStr: "",
				argsLen: 0,
			},
			want:    []byte("+PONG\r\n"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &command{
				cmd:     tt.fields.cmd,
				argsStr: tt.fields.argsStr,
				argsLen: tt.fields.argsLen,
			}
			got, err := c.Response()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response() got = %v, want %v", got, tt.want)
			}
		})
	}
}
