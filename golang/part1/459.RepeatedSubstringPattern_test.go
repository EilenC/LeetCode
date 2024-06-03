package part1

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "item1",
			args: args{
				s: "abab",
			},
			want: true,
		},
		{
			name: "item2",
			args: args{
				s: "aba",
			},
			want: false,
		},
		{
			name: "item3",
			args: args{
				s: "abcabcabcabc",
			},
			want: true,
		},
		{
			name: "item4",
			args: args{
				s: "abac",
			},
			want: false,
		},
		{
			name: "item5",
			args: args{
				s: "abcabc",
			},
			want: true,
		},
		{
			name: "item6",
			args: args{
				s: "bb",
			},
			want: true,
		},
		{
			name: "item7",
			args: args{
				s: "abaababaab",
			},
			want: true,
		},
		{
			name: "item8",
			args: args{
				s: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
