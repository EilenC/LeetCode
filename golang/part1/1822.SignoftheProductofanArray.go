package part1

//There is a function signFunc(x) that returns:
//
//1 if x is positive.
//-1 if x is negative.
//0 if x is equal to 0.
//You are given an integer array nums. Let product be the product of all values in the array nums.
//
//Return signFunc(product).
//
//
//
//Example 1:
//
//Input: nums = [-1,-2,-3,-4,3,2,1]
//Output: 1
//Explanation: The product of all values in the array is 144, and signFunc(144) = 1
//Example 2:
//
//Input: nums = [1,5,0,2,-3]
//Output: 0
//Explanation: The product of all values in the array is 0, and signFunc(0) = 0
//Example 3:
//
//Input: nums = [-1,1,-1,1,-1]
//Output: -1
//Explanation: The product of all values in the array is -1, and signFunc(-1) = -1

/*
*
1. 循环判断nums中 负数数量, 0 数量
2. 有0 则直接返回 0
3. 负数 0个或者偶数个 直接返回1
4. 负数单个返回 -1
*/
func arraySign(nums []int) int {
	n := 0
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			n++
		}
	}
	if n%2 == 0 {
		return 1
	}
	return -1
}
