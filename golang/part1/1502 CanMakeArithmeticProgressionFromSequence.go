package part1

/**
给你一个数字数组 arr 。

如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。

如果可以重新排列数组形成等差数列，请返回 true ；否则，返回 false 。



示例 1：

输入：arr = [3,5,1]
输出：true
解释：对数组重新排序得到 [1,3,5] 或者 [5,3,1] ，任意相邻两项的差分别为 2 或 -2 ，可以形成等差数列。
示例 2：

输入：arr = [1,2,4]
输出：false
解释：无法通过重新排序得到等差数列。


提示：

2 <= arr.length <= 1000
-10^6 <= arr[i] <= 10^6
*/

/*
*
1. 将arr排序,然后遍历计算每个数字差值
*/
//func canMakeArithmeticProgression(arr []int) bool {
//	sort.Ints(arr)
//	p := arr[1] - arr[0]
//	length := len(arr)
//	for i := 1; i < length-1; i++ {
//		if arr[i+1]-arr[i] != p {
//			return false
//		}
//	}
//	return true
//}

/*
*
1. 先取出arr中最大值与最小值,这样max-min就能计算出公差
2. 如果是等差数列每个数据arr[i] = min + i*公差
*/

func partition(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}

func canMakeArithmeticProgression(arr []int) bool {
	length := len(arr)
	QuickSort(arr, 0, length-1)
	p := arr[1] - arr[0]
	for i := 1; i < length-1; i++ {
		if arr[i+1]-arr[i] != p {
			return false
		}
	}
	return true
}
