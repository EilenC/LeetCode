package part1

/**
你一个整数数组 nums ，如果 nums 至少 包含 2 个元素，你可以执行以下操作：

选择 nums 中的前两个元素并将它们删除。
一次操作的 分数 是被删除元素的和。

在确保 所有操作分数相同 的前提下，请你求出 最多 能进行多少次操作。

请你返回按照上述要求 最多 可以进行的操作次数。



示例 1：

输入：nums = [3,2,1,4,5]
输出：2
解释：我们执行以下操作：
- 删除前两个元素，分数为 3 + 2 = 5 ，nums = [1,4,5] 。
- 删除前两个元素，分数为 1 + 4 = 5 ，nums = [5] 。
由于只剩下 1 个元素，我们无法继续进行任何操作。
示例 2：

输入：nums = [3,2,6,1,4]
输出：1
解释：我们执行以下操作：
- 删除前两个元素，分数为 3 + 2 = 5 ，nums = [6,1,4] 。
由于下一次操作的分数与前一次不相等，我们无法继续进行任何操作。


提示：

2 <= nums.length <= 100
1 <= nums[i] <= 1000
*/

/*
*
1. 根据题目意义,若当前[]int小于2,则返回0
2. 若长度是2或者3 则返回1
3. 若长度超过3则开始计算
4. 最后把循环的i/2取整即可得出步数
*/
func maxOperations(nums []int) int {
	length := len(nums)
	if length < 2 {
		return 0
	}
	if length == 2 || length == 3 {
		return 1
	}
	nums[0] += nums[1] //节省空间,将nums[0]当做最开始两项的和
	nums[1] = 3        //节省空间,将nums[1]作为for循环遍历的下标,方便计算步数
	for ; nums[1] < length; nums[1] += 2 {
		if nums[nums[1]]+nums[nums[1]-1] != nums[0] {
			break
		}
	}
	return nums[1] / 2
}
