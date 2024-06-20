package part1

import "testing"

func Test_countBeautifulPairs(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：nums = [2,5,1,4]",
			args: args{
				nums: []int{2, 5, 1, 4},
			},
			want: 5,
		},
		{
			name: "输入：nums = [11,21,12]",
			args: args{
				nums: []int{11, 21, 12},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBeautifulPairs(tt.args.nums); got != tt.want {
				t.Errorf("countBeautifulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
