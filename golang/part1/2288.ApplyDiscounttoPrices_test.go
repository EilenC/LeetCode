package part1

import "testing"

func Test_discountPrices(t *testing.T) {
	type args struct {
		sentence string
		discount int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "输入：sentence = there are $1 $2 and 5$ candies in the shop, discount = 50",
			args: args{
				sentence: "there are $1 $2 and 5$ candies in the shop",
				discount: 50,
			},
			want: "there are $0.50 $1.00 and 5$ candies in the shop",
		},
		{
			name: "输入：sentence = 1 2 $3 4 $5 $6 7 8$ $9 $10$, discount = 100",
			args: args{
				sentence: "1 2 $3 4 $5 $6 7 8$ $9 $10$",
				discount: 100,
			},
			want: "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$",
		},
		{
			name: "输入：sentence = $76111 ab $6 $, discount = 48",
			args: args{
				sentence: "$76111 ab $6 $",
				discount: 48,
			},
			want: "$39577.72 ab $3.12 $",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := discountPrices(tt.args.sentence, tt.args.discount)
			if got != tt.want {
				t.Errorf("discountPrices() = %v, want %v", got, tt.want)
			}
		})
	}
}
