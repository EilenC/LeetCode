package part1

import "testing"

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "item1",
			args: args{
				moves: "UD",
			},
			want: true,
		},
		{
			name: "item2",
			args: args{
				moves: "LL",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := judgeCircle(tt.args.moves); got != tt.want {
				t.Errorf("judgeCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}
