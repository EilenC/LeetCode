package part1

import "testing"

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "输入: s = A man, a plan, a canal: Panama",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "输入：s = race a car",
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "输入：s =  ",
			args: args{
				s: " ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}
