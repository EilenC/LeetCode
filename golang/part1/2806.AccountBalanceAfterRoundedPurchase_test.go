package part1

import "testing"

func Test_accountBalanceAfterPurchase(t *testing.T) {
	type args struct {
		purchaseAmount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "输入：purchaseAmount = 9",
			args: args{
				purchaseAmount: 9,
			},
			want: 90,
		},
		{
			name: "输入：purchaseAmount = 15",
			args: args{
				purchaseAmount: 15,
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accountBalanceAfterPurchase(tt.args.purchaseAmount); got != tt.want {
				t.Errorf("accountBalanceAfterPurchase() = %v, want %v", got, tt.want)
			}
		})
	}
}
