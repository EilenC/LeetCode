package part1

import "testing"

func Test_removeTrailingZeros(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "输入：num = 51230100",
			args: args{
				num: "51230100",
			},
			want: "512301",
		},
		{
			name: "输入：num = 123",
			args: args{
				num: "123",
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeTrailingZeros(tt.args.num); got != tt.want {
				t.Errorf("removeTrailingZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
