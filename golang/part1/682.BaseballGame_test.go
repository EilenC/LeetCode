package part1

import "testing"

func Test_calPoints(t *testing.T) {
	type args struct {
		operations []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "item1",
			args: args{
				operations: []string{"5", "2", "C", "D", "+"},
			},
			want: 30,
		},
		{
			name: "item2",
			args: args{
				operations: []string{"5", "-2", "4", "C", "D", "9", "+", "+"},
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPoints(tt.args.operations); got != tt.want {
				t.Errorf("calPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
