package part1

import "testing"

func Test_canMakeArithmeticProgression(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "arr = [3,5,1]",
			args: args{
				arr: []int{3, 5, 1},
			},
			want: true,
		},
		{
			name: "arr = [1,2,4]",
			args: args{
				arr: []int{1, 2, 4},
			},
			want: false,
		},
		{
			name: "arr = [1,4,4,7]",
			args: args{
				arr: []int{1, 4, 4, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.args.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
