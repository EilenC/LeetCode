package part1

import "testing"

func Test_tictactoe(t *testing.T) {
	type args struct {
		moves [][]int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item1",
			args: args{
				moves: [][]int{[]int{0, 0}, []int{2, 0}, []int{1, 1}, []int{2, 1}, []int{2, 2}},
			},
			want: "A",
		},
		{
			name: "item2",
			args: args{
				moves: [][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}},
			},
			want: "B",
		},
		{
			name: "item3",
			args: args{
				moves: [][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}},
			},
			want: "Draw",
		},
		{
			name: "item4",
			args: args{
				moves: [][]int{{0, 0}, {1, 1}},
			},
			want: "Pending",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tictactoe(tt.args.moves); got != tt.want {
				t.Errorf("tictactoe() = %v, want %v", got, tt.want)
			}
		})
	}
}
