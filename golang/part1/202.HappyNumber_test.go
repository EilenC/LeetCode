package part1

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "输入：n = 19",
			args: args{
				n: 19,
			},
			want: true,
		},
		{
			name: "输入：n = 2",
			args: args{
				n: 2,
			},
			want: false,
		},
		{
			name: "输入：n = 7",
			args: args{
				n: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
