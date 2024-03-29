package part1

import "testing"

func Test_diagonalSum(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "item1",
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSum(tt.args.mat); got != tt.want {
				t.Errorf("diagonalSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
