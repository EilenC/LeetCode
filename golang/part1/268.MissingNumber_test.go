package part1

import "testing"

func Test_missingNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,0,1]",
			args: args{
				nums: []int{3, 0, 1},
			},
			want: 2,
		},
		{
			name: "[0,1]",
			args: args{
				nums: []int{0, 1},
			},
			want: 2,
		},
		{
			name: "[9,6,4,2,3,5,7,0,1]",
			args: args{
				nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			want: 8,
		},
		{
			name: "[0]",
			args: args{
				nums: []int{0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumber(tt.args.nums); got != tt.want {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
