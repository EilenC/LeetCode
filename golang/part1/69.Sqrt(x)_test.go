package part1

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：x = 4",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "输入：x = 8",
			args: args{
				x: 8,
			},
			want: 2,
		},
		{
			name: "输入：x = 9",
			args: args{
				x: 9,
			},
			want: 3,
		},
		{
			name: "输入：x = 1",
			args: args{
				x: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
