package part1

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入: a = \"aba\", b = \"cdc\"",
			args: args{
				a: "aba",
				b: "cdc",
			},
			want: 3,
		},
		{
			name: "输入：a = \"aaa\", b = \"bbb\"",
			args: args{
				a: "aaa",
				b: "bbb",
			},
			want: 3,
		},
		{
			name: "输入：a = \"aaa\", b = \"aaa\"",
			args: args{
				a: "aaa",
				b: "aaa",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
