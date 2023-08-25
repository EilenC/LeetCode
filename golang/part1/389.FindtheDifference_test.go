package part1

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "item1",
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: 'e',
		},
		{
			name: "item2",
			args: args{
				s: "",
				t: "y",
			},
			want: 'y',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
