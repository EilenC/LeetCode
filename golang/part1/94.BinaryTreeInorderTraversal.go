package part1

/**
给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[1,3,2]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100


进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodeItem struct {
	View bool
	Node *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var (
		res      []int
		tempItem NodeItem
	)

	stack := []NodeItem{{
		View: true,
		Node: root,
	}}
	length := len(stack)
	for length != 0 {
		tempItem = stack[length-1]
		stack = stack[:length-1]
		if tempItem.Node == nil {
			length = len(stack)
			continue
		}
		if tempItem.View {
			stack = append(stack, NodeItem{
				View: true,
				Node: tempItem.Node.Right,
			})
			stack = append(stack, NodeItem{
				Node: tempItem.Node,
			})
			stack = append(stack, NodeItem{
				View: true,
				Node: tempItem.Node.Left,
			})
		} else {
			res = append(res, tempItem.Node.Val)
		}
		length = len(stack)
	}
	return res
}
