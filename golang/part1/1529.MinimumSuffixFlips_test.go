package part1

import "testing"

func Test_minFlips(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：target = 10111",
			args: args{
				target: "10111",
			},
			want: 3,
		},
		{
			name: "输入：target = 101",
			args: args{
				target: "101",
			},
			want: 3,
		},
		{
			name: "输入：target = 00000",
			args: args{
				target: "00000",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.target); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
