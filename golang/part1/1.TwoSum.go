package part1

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]
*/

/*
*
1. 遍历数组,按顺序从头开始遍历数字,然后再遍历从数字开始后额度所有数字直到结束
*/
func twoSum(nums []int, target int) []int {
	length := len(nums)
	for i, num := range nums {
		ans := target - num
		for j := i + 1; j < length; j++ {
			if nums[j] == ans {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
