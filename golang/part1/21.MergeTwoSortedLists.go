package part1

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	nl := new(ListNode)
	top := nl
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			top.Next = list2
			list2 = list2.Next
		} else {
			top.Next = list1
			list1 = list1.Next
		}
		top = top.Next
	}
	if list1 != nil {
		top.Next = list1
	}
	if list2 != nil {
		top.Next = list2
	}
	return nl.Next
}
