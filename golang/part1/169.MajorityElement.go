package part1

/**
给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。



示例 1：

输入：nums = [3,2,3]
输出：3
示例 2：

输入：nums = [2,2,1,1,1,2,2]
输出：2
*/

/*
*
1. 遍历数组并且用map存储每个数字出现的次数
2. 遍历结果的map,返回count最多的数字
*/
func majorityElement(nums []int) int {
	aMap := make(map[int]int)
	for _, num := range nums {
		aMap[num] = aMap[num] + 1
	}
	maxNum := 0
	maxCount := 0
	for num, count := range aMap {
		if count > maxCount {
			maxNum = num
			maxCount = count
		}
	}
	return maxNum
}
