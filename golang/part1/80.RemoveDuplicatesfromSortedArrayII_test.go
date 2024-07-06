package part1

import "testing"

func Test_removeDuplicates2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：nums = [1,1,1,2,2,3]",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "输入：nums = [0,0,1,1,1,1,2,3,3]",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates2(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}
