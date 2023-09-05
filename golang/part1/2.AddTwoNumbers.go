package part1

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//
//你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.

/**
* Definition for singly-linked list.
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	var current = dummy
	var suffix, sum int

	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		total := n1 + n2
		if suffix > 0 {
			total += suffix
		}
		if total > 9 {
			sum, suffix = total%10, total/10
		} else {
			sum = total
			suffix = 0
		}

		// 创建一个新的链表节点，并将其添加到当前节点的后面
		current.Next = &ListNode{Val: sum}
		current = current.Next
		sum = 0
	}

	// 如果最后还有进位，创建一个额外的节点
	if suffix > 0 {
		current.Next = &ListNode{Val: suffix}
	}

	return dummy.Next
}
