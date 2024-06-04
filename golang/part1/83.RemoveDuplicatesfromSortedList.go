package part1

/**
给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。



示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]


提示：

链表中节点数目在范围 [0, 300] 内
-100 <= Node.val <= 100
题目数据保证链表已经按升序 排列
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
*
1. 由于链表已经排过序,所以每遇到相同的Node则都将Next指向Next的Next
2. 最后返回原list即可
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	next := head
	for next.Next != nil {
		if next.Val == next.Next.Val {
			next.Next = next.Next.Next
		} else {
			next = next.Next
		}
	}
	return head
}
