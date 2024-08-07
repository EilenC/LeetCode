package part1

import "testing"

func Test_maxProfit2(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：prices = [7,1,5,3,6,4]",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 7,
		},
		{
			name: "输入：prices = [1,2,3,4,5]",
			args: args{
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "输入：prices = [7,6,4,3,1]",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
