package part1

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
//请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//
//
//示例 1:
//
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]
//示例 2:
//
//输入: nums = [0]
//输出: [0]
//
//
//提示:
//
//1 <= nums.length <= 104
//-231 <= nums[i] <= 231 - 1
//
//
//进阶：你能尽量减少完成的操作次数吗？
//func moveZeroes(nums []int) {
//	var (
//		zeroNum int
//	)
//	newArr := nums[0:]
//	nums = nums[:0]
//	for _, v := range newArr {
//		if v == 0 {
//			zeroNum++
//			continue
//		}
//		nums = append(nums, int(v))
//	}
//	for i := 0; i < zeroNum; i++ {
//		nums = append(nums, 0)
//	}
//	return
//}

func moveZeroes(nums []int) {
	var i, j int
	for _, _ = range nums {
		if nums[j] != 0 {
			if j > i {
				nums[i] = nums[j]
				nums[j] = 0
			}
			//nums[j], nums[i] = nums[i], nums[j]
			i++
		}
		j++
	}
}
