package part1

import "testing"

func Test_toHex(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "item1",
			args: args{
				num: 26,
			},
			want: "1a",
		},
		{
			name: "item2",
			args: args{
				num: -1,
			},
			want: "ffffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.args.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
