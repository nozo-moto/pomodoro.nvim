package command

import "testing"

func Test_getFormatedNowTime(t *testing.T) {
	type args struct {
		nowTime int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormatedNowTime(tt.args.nowTime); got != tt.want {
				t.Errorf("getFormatedNowTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
