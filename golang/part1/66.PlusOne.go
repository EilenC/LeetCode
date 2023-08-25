package part1

//给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//
//
//示例 1：
//
//输入：digits = [1,2,3]
//输出：[1,2,4]
//解释：输入数组表示数字 123。
//示例 2：
//
//输入：digits = [4,3,2,1]
//输出：[4,3,2,2]
//解释：输入数组表示数字 4321。
//示例 3：
//
//输入：digits = [0]
//输出：[1]

/*
*
1.从尾数循环遍历,逐步+1
2.判断余数与进位
3.最后一位数依然有进位,则append slice
4.检查是否有扩容slice 将slice顺序调整
*/
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				digits = append(digits, 1)
			}
			continue
		}
		digits[i] = digits[i] + 1
		break
	}
	if n != len(digits) {
		res := digits[n:]
		res = append(res, digits[:n]...)
		return res
	}
	return digits
}
