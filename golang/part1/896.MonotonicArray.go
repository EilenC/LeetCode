package part1

//如果数组是单调递增或单调递减的，那么它是 单调 的。
//
//如果对于所有 i <= j，nums[i] <= nums[j]，那么数组 nums 是单调递增的。 如果对于所有 i <= j，nums[i]> = nums[j]，那么数组 nums 是单调递减的。
//
//当给定的数组 nums 是单调数组时返回 true，否则返回 false。

func isMonotonic(nums []int) bool {
	flag1, flag2 := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			flag1 = false
		}
		if nums[i+1] < nums[i] {
			flag2 = false
		}
	}
	return flag1 || flag2
}
