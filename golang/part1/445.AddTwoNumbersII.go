package part1

import "sync"

/**
给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。



示例1：



输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]
示例2：

输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]
示例3：

输入：l1 = [0], l2 = [0]
输出：[0]


提示：

链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
1. 使用栈进行计算与答案记录
*/
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	//遍历两个链表,分别压入两个栈
	s1, s2, ans := NewStack(), NewStack(), NewStack()
	for l1 != nil || l2 != nil {
		if l1 != nil {
			s1.Push(l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			s2.Push(l2.Val)
			l2 = l2.Next
		}
	}
	var dummy = &ListNode{}
	var current = dummy
	var suffix, sum int
	var n1, n2 interface{}
	for s1.Peek() != nil || s2.Peek() != nil {
		n1, n2 = s1.Pop(), s2.Pop()
		if n1 == nil {
			n1 = 0
		}
		if n2 == nil {
			n2 = 0
		}
		total := n1.(int) + n2.(int)
		if suffix > 0 {
			total += suffix
		}
		if total > 9 {
			sum, suffix = total%10, total/10
		} else {
			sum = total
			suffix = 0
		}
		//答案也是压入栈
		ans.Push(sum)
		sum = 0
	}
	// 如果最后还有进位，创建一个额外的节点
	if suffix > 0 {
		ans.Push(suffix)
	}
	//出栈拼接链表
	for ans.Peek() != nil {
		n1 = ans.Pop()
		// 创建一个新的链表节点，并将其添加到当前节点的后面
		current.Next = &ListNode{Val: n1.(int)}
		current = current.Next
	}
	return dummy.Next
}

type (
	Stack struct {
		top    *node
		length int
		lock   *sync.RWMutex
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0, &sync.RWMutex{}}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	n := &node{value, this.top}
	this.top = n
	this.length++
}
