package part1

import "testing"

func Test_findMiddleIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：nums = [2,3,-1,8,4]",
			args: args{
				nums: []int{2, 3, -1, 8, 4},
			},
			want: 3,
		},
		{
			name: "输入：nums = [1,-1,4]",
			args: args{
				nums: []int{1, -1, 4},
			},
			want: 2,
		},
		{
			name: "输入：nums = [2,5]",
			args: args{
				nums: []int{2, 5},
			},
			want: -1,
		},
		{
			name: "输入：nums = [1]",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "输入：nums = [1,1]",
			args: args{
				nums: []int{1, 1},
			},
			want: -1,
		},
		{
			name: "输入：nums = [-1,2]",
			args: args{
				nums: []int{-1, 2},
			},
			want: -1,
		},
		{
			name: "输入：nums = [4,0]",
			args: args{
				nums: []int{4, 0},
			},
			want: 0,
		},
		{
			name: "输入：nums = [0,0,0,0]",
			args: args{
				nums: []int{0, 0, 0, 0},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMiddleIndex(tt.args.nums); got != tt.want {
				t.Errorf("findMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
