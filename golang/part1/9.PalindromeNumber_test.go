package part1

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "输入：x = 121",
			args: args{
				x: 121,
			},
			want: true,
		},
		{
			name: "输入：x = -121",
			args: args{
				x: -121,
			},
			want: false,
		},
		{
			name: "输入：x = 10",
			args: args{
				x: 10,
			},
			want: false,
		},
		{
			name: "输入：x = 123454321",
			args: args{
				x: 123454321,
			},
			want: true,
		},
		{
			name: "输入：x = 1234554321",
			args: args{
				x: 1234554321,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
