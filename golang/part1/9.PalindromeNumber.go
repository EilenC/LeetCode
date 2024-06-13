package part1

import "strconv"

/**
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数
是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。


示例 1：

输入：x = 121
输出：true
示例 2：

输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3：

输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。


提示：

-231 <= x <= 231 - 1
*/

/*
*
1. 取出右边第一位数字算法是 x % 10 然后 x = x / 10
2. 取出左边第一位数字算法是 x
*/
//func isPalindrome(x int) bool {
//	// 特殊情况：
//	// 如上所述，当 x < 0 时，x 不是回文数。
//	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
//	// 则其第一位数字也应该是 0
//	// 只有 0 满足这一属性
//	if x < 0 || (x%10 == 0 && x != 0) {
//		return false
//	}
//	var revertedNumber int
//
//	for !(x <= revertedNumber) {
//		revertedNumber = (revertedNumber * 10) + (x % 10)
//		x /= 10
//	}
//	return x == revertedNumber || x == revertedNumber/10
//}

/*
*
1. 将数字转换为字符串
2. 使用双下标一一对比,i++,j-- 当j <= i 时推出判断
*/
func isPalindrome(x int) bool {
	xs := strconv.Itoa(x)
	j := len(xs) - 1
	x = 0

	for j > x {
		if xs[x] != xs[j] {
			return false
		}
		j--
		x++
	}
	return true
}
