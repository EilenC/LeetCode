package part1

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：nums = [2,1,2]",
			args: args{
				nums: []int{2, 1, 2},
			},
			want: 5,
		},
		{
			name: "输入：nums = [1,2,1,10]",
			args: args{
				nums: []int{1, 2, 1, 10},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.nums); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
