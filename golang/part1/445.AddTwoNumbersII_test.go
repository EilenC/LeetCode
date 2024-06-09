package part1

import (
	"testing"
)

func Test_addTwoNumbers2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "输入：l1 = [7,2,4,3], l2 = [5,6,4]",
			args: args{
				l1: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  7,
							Next: nil,
						},
					},
				},
			},
		},
		{
			name: "输入：l1 = [2,4,3], l2 = [5,6,4]",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val:  7,
						Next: nil,
					},
				},
			},
		},
		{
			name: "输入：l1 = [0], l2 = [0]",
			args: args{
				l1: &ListNode{
					Val:  0,
					Next: nil,
				},
				l2: &ListNode{
					Val:  0,
					Next: nil,
				},
			},
			want: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers2(tt.args.l1, tt.args.l2)
			if got == nil {
				t.Errorf("addTwoNumbers2() = %v, want %v", got, tt.want)
			}
			for got != nil || tt.want != nil {
				n1 := got.Val
				n2 := tt.want.Val

				if n1 != n2 {
					t.Errorf("addTwoNumbers2() = %v, want %v", got, tt.want)
				}
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}
