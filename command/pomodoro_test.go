package command

import (
	"testing"
	"time"
)

func Test_getFormatedNowTime(t *testing.T) {
	type args struct {
		nowTime int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				nowTime: int(time.Minute.Seconds() * 25),
			},
			want: "25:0",
		},
		{
			name: "test",
			args: args{
				nowTime: int(time.Minute.Seconds()*25) + int(time.Second.Seconds()*25),
			},
			want: "25:25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormatedNowTime(tt.args.nowTime); got != tt.want {
				t.Errorf("getFormatedNowTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
