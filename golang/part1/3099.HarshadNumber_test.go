package part1

import "testing"

func Test_sumOfTheDigitsOfHarshadNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入： x = 18",
			args: args{
				x: 18,
			},
			want: 9,
		},
		{
			name: "输入： x = 23",
			args: args{
				x: 23,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfTheDigitsOfHarshadNumber(tt.args.x); got != tt.want {
				t.Errorf("sumOfTheDigitsOfHarshadNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
