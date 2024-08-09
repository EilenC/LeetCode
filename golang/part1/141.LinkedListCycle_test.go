package part1

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	args_1_3 := &ListNode{Val: 3}
	args_1_2 := &ListNode{Val: 2}
	args_1_0 := &ListNode{Val: 0}
	args_1_4 := &ListNode{Val: 4}
	args_1_3.Next = args_1_2
	args_1_2.Next = args_1_0
	args_1_0.Next = args_1_4
	args_1_4.Next = args_1_2

	args_2_1 := &ListNode{Val: 1}
	args_2_2 := &ListNode{Val: 2}

	args_2_1.Next = args_2_2
	args_2_2.Next = args_2_1

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: `输入：head = [3,2,0,-4], pos = 1`,
			args: args{
				head: args_1_3,
			},
			want: true,
		},
		{
			name: `输入：head = [1,2], pos = 0`,
			args: args{
				head: args_2_1,
			},
			want: true,
		},
		{
			name: `输入：head = [1], pos = -1`,
			args: args{
				head: &ListNode{Val: 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
