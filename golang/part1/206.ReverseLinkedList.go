package part1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var top *ListNode
	for x := head; x != nil; x = x.Next {
		top = &ListNode{Val: x.Val, Next: top}
	}
	return top
}
