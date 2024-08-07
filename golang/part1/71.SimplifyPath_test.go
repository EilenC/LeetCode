package part1

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `输入：path = "/home/"`,
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: `输入：path = "/../"`,
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: `输入：path = "/home//foo/"`,
			args: args{
				path: "/home//foo",
			},
			want: "/home/foo",
		},
		{
			name: `输入：path = "/a/./b/../../c/"`,
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
