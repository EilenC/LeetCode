package part1

/*
*
给你一个整数数组 nums。

返回两个（不一定不同的）质数在 nums 中 下标 的 最大距离。

示例 1：

输入： nums = [4,2,9,5,3]

输出： 3

解释： nums[1]、nums[3] 和 nums[4] 是质数。因此答案是 |4 - 1| = 3。

示例 2：

输入： nums = [4,8,2,8]

输出： 0

解释： nums[2] 是质数。因为只有一个质数，所以答案是 |2 - 2| = 0。

提示：

1 <= nums.length <= 3 * 105
1 <= nums[i] <= 100
输入保证 nums 中至少有一个质数。
*/
func maximumPrimeDifference(nums []int) int {
	var i, j int
	for !isPrime(nums[i]) {
		i++
	}
	j = len(nums) - 1
	for !isPrime(nums[j]) {
		j--
	}
	return j - i
}

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
