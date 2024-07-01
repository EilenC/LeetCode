package part1

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "输入：root = [1,null,2,3]",
			args: args{
				root: &TreeNode{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
				},
			},
			want: []int{3, 2, 1},
		},
		{
			name: "输入：root = []",
			args: args{},
			want: nil,
		},
		{
			name: "输入：root = [1]",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := postorderTraversal(tt.args.root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
