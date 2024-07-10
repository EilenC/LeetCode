package part1

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "输入：num = 3749",
			args: args{
				num: 3749,
			},
			want: "MMMDCCXLIX",
		},
		{
			name: "输入：num = 58",
			args: args{
				num: 58,
			},
			want: "LVIII",
		},
		{
			name: "输入：num = 1994",
			args: args{
				num: 1994,
			},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
