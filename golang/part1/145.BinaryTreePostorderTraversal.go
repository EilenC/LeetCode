package part1

/**
给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。



示例 1：


输入：root = [1,null,2,3]
输出：[3,2,1]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]


提示：

树中节点的数目在范围 [0, 100] 内
-100 <= Node.val <= 100

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
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
				Node: tempItem.Node,
			})
			stack = append(stack, NodeItem{
				View: true,
				Node: tempItem.Node.Right,
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
