package part1

import "testing"

func Test_trap(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			wantAns: 6,
		},
		{
			name: "输入：height = [4,2,0,3,2,5]",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			wantAns: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := trap(tt.args.height); gotAns != tt.wantAns {
				t.Errorf("trap() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
