package part1

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "item1",
			args: args{
				list1: new(ListNode),
				list2: new(ListNode),
			},
			want: new(ListNode),
		},
	}
	for _, tt := range tests {
		tt.args.list1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
		tt.args.list2 = &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
		tt.want = &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}

		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			for got != nil || tt.want != nil {
				if got.Val != tt.want.Val {
					t.Errorf("mergeTwoLists() = %v, want %v", got.Val, tt.want.Val)
					break
				} else {
					got = got.Next
					tt.want = tt.want.Next
				}
			}
		})
	}
}
