package part1

import "testing"

func Test_maxOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：nums = [3,2,1,4,5]",
			args: args{
				nums: []int{3, 2, 1, 4, 5},
			},
			want: 2,
		},
		{
			name: "输入：nums = [3,2,6,1,4]",
			args: args{
				nums: []int{3, 2, 6, 1, 4},
			},
			want: 1,
		},
		{
			name: "输入：nums = [1]",
			args: args{
				nums: []int{1},
			},
			want: 0,
		},
		{
			name: "输入：nums = [1,2]",
			args: args{
				nums: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "输入：nums = [1,2,3]",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxOperations(tt.args.nums); got != tt.want {
				t.Errorf("maxOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
