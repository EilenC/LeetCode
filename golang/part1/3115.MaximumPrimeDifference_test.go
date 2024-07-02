package part1

import "testing"

func Test_maximumPrimeDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入： nums = [4,2,9,5,3]",
			args: args{
				nums: []int{4, 2, 9, 5, 3},
			},
			want: 3,
		},
		{
			name: "输入： nums = [4,8,2,8]",
			args: args{
				nums: []int{4, 8, 2, 8},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumPrimeDifference(tt.args.nums); got != tt.want {
				t.Errorf("maximumPrimeDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
