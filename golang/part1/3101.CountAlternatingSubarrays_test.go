package part1

import "testing"

func Test_countAlternatingSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "输入： nums = [0,1,1,1]",
			args: args{
				nums: []int{0, 0, 1, 1},
			},
			want: 5,
		},
		{
			name: "输入： nums = [1,0,1,0]",
			args: args{
				nums: []int{1, 0, 1, 0},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAlternatingSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("countAlternatingSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
