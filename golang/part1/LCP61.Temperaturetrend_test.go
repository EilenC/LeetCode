package part1

import "testing"

func Test_temperatureTrend(t *testing.T) {
	type args struct {
		temperatureA []int
		temperatureB []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入： temperatureA = [21,18,18,18,31] temperatureB = [34,32,16,16,17]",
			args: args{
				temperatureA: []int{21, 18, 18, 18, 31},
				temperatureB: []int{34, 32, 16, 16, 17},
			},
			want: 2,
		},
		{
			name: "输入： temperatureA = [5,10,16,-6,15,11,3] temperatureB = [16,22,23,23,25,3,-16]",
			args: args{
				temperatureA: []int{5, 10, 16, -6, 15, 11, 3},
				temperatureB: []int{16, 22, 23, 23, 25, 3, -16},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := temperatureTrend(tt.args.temperatureA, tt.args.temperatureB); got != tt.want {
				t.Errorf("temperatureTrend() = %v, want %v", got, tt.want)
			}
		})
	}
}
