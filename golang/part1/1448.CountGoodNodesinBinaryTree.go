package part1

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//给你一棵根为 root 的二叉树，请你返回二叉树中好节点的数目。
//
//「好节点」X 定义为：从根到该节点 X 所经过的节点中，没有任何节点的值大于 X 的值。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, maxNum int) int {
	if root == nil {
		return 0
	}
	res := 0

	if root.Val >= maxNum {
		res++
		maxNum = root.Val
	}

	res += dfs(root.Left, maxNum) + dfs(root.Right, maxNum)
	return res
}

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt64)
}
