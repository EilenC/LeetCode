package part1

import (
	"reflect"
	"testing"
)

func Test_numberGame(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "输入：nums = [5,4,2,3]",
			args: args{
				nums: []int{5, 4, 2, 3},
			},
			want: []int{3, 2, 5, 4},
		},
		{
			name: "输入：nums = [2,5]",
			args: args{
				nums: []int{2, 5},
			},
			want: []int{5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberGame(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
