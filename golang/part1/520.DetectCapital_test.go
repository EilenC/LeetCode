package part1

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "输入：word = USA",
			args: args{
				word: "USA",
			},
			want: true,
		},
		{
			name: "输入：word = FlaG",
			args: args{
				word: "FlaG",
			},
			want: false,
		},
		{
			name: "输入：word = Leetcode",
			args: args{
				word: "Leetcode",
			},
			want: true,
		},
		{
			name: "输入：word = mL",
			args: args{
				word: "mL",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
