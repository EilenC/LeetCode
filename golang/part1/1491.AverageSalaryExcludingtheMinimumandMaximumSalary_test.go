package part1

import "testing"

func Test_average(t *testing.T) {
	type args struct {
		salary []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "item1",
			args: args{
				salary: []int{4000, 3000, 1000, 2000},
			},
			want: 2500.00000,
		},
		{
			name: "item1",
			args: args{
				salary: []int{1000, 2000, 3000},
			},
			want: 2000.00000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.salary); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
