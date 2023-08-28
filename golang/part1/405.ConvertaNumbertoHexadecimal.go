package part1

//Given an integer num, return a string representing its hexadecimal representation. For negative integers, two’s complement method is used.
//
//All the letters in the answer string should be lowercase characters, and there should not be any leading zeros in the answer except for the zero itself.
//
//Note: You are not allowed to use any built-in library method to directly solve this problem.
//
//
//
//Example 1:
//
//Input: num = 26
//Output: "1a"
//Example 2:
//
//Input: num = -1
//Output: "ffffffff"
//
//
//Constraints:
//
//-231 <= num <= 231 - 1

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	ans := ""
	for num != 0 && len(ans) < 8 {
		ans = string(hex[byte(num&0xf)]) + ans
		num >>= 4
	}
	return ans
}
