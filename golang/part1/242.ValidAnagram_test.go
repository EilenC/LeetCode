package part1

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "item1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "item2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "item3",
			args: args{
				s: "ac",
				t: "bb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
