package part1

//给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。

/**
示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。

*/

/*
*
1. 计算数组(slice)长度,使用一个[]bool rows 来表示所有包含的数字范围
2. 循环已有的 nums 数据,在有的数据上将bool改为true
3. 查询 rows 中值为 false 的 index即可
*/
func missingNumber(nums []int) int {
	n := len(nums)
	rows := make([]bool, n+1)
	for _, num := range nums {
		rows[num] = true
	}
	for number, row := range rows {
		if row == false {
			return number
		}
	}
	return 0
}

/**
1. 利用异或的特性,一个数与自己异或0,与0异或等于自身
2. 把接收到的数据与整个rows长度异或一遍最后再补齐+1的闭区间数字
*/
//func missingNumber(nums []int) int {
//	ans := 0
//	for i, num := range nums {
//		ans ^= i ^ num
//	}
//	return ans ^ len(nums)
//}
