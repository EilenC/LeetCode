package part1

import "testing"

func Test_lemonadeChange(t *testing.T) {
	type args struct {
		bills []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "bills = [5,5,5,10,20]",
			args: args{
				bills: []int{5, 5, 5, 10, 20},
			},
			want: true,
		},
		{
			name: "bills = [5,5,10,10,20]",
			args: args{
				bills: []int{5, 5, 10, 10, 20},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lemonadeChange(tt.args.bills); got != tt.want {
				t.Errorf("lemonadeChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
