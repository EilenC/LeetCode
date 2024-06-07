package part1

import "strings"

/**
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。



示例 1：

输入:a = "11", b = "1"
输出："100"
示例 2：

输入：a = "1010", b = "1011"
输出："10101"


提示：

1 <= a.length, b.length <= 104
a 和 b 仅由字符 '0' 或 '1' 组成
字符串如果不是 "0" ，就不含前导零
*/

/*
*
1. 通过计算长度,将两个字符串长度改为一致,较短的部分左边补0
2. 反向遍历模拟二进制计算进位等
*/
func binaryCalcPlus(a, b int8, plus bool) (int8, bool) {
	if a == '0' && b == '0' {
		if plus {
			return '1', false
		}
		return '0', false
	}
	if a == '1' && b == '0' || a == '0' && b == '1' {
		if plus {
			return '0', true
		}
		return '1', false
	}
	if a == '1' && b == '1' {
		if plus {
			return '1', true
		}
		return '0', true
	}
	return 0, false
}

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	var ans []int8
	var av int8
	var plus bool
	if la < lb {
		for i := 0; i < lb-la; i++ {
			a = "0" + a
		}
		ans = make([]int8, 0, lb)
		for i := lb - 1; i >= 0; i-- {
			av, plus = binaryCalcPlus(int8(a[i]), int8(b[i]), plus)
			ans = append(ans, av)
		}
	} else {
		for i := 0; i < la-lb; i++ {
			b = "0" + b
		}
		ans = make([]int8, 0, la)
		for i := la - 1; i >= 0; i-- {
			av, plus = binaryCalcPlus(int8(b[i]), int8(a[i]), plus)
			ans = append(ans, av)
		}
	}
	if plus {
		ans = append(ans, '1')
	}
	vs := strings.Builder{}
	for i := len(ans) - 1; i >= 0; i-- {
		vs.WriteByte(byte(ans[i]))
	}
	return vs.String()
}
