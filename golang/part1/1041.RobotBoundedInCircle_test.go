package part1

import "testing"

func Test_isRobotBounded(t *testing.T) {
	type args struct {
		instructions string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "item1",
			args: args{
				instructions: "GGLLGG",
			},
			want: true,
		},
		{
			name: "item1",
			args: args{
				instructions: "GG",
			},
			want: false,
		},
		{
			name: "item1",
			args: args{
				instructions: "GL",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRobotBounded(tt.args.instructions); got != tt.want {
				t.Errorf("isRobotBounded() = %v, want %v", got, tt.want)
			}
		})
	}
}
