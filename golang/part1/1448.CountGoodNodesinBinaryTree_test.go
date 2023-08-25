package part1

import "testing"

func Test_goodNodes(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "item1",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 5},
					},
				},
			},
			want: 4,
		},
		{
			name: "item2",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   2,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodNodes(tt.args.root); got != tt.want {
				t.Errorf("goodNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
