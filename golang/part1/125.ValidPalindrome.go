package part1

import "fmt"

/**
如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。

字母和数字都属于字母数字字符。

给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。



示例 1：

输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
示例 2：

输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
示例 3：

输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。


提示：

1 <= s.length <= 2 * 105
s 仅由可打印的 ASCII 字符组成
*/

/*
*
1. 头尾指针相互移动,遇到正常的字符串才移动(排除空格之类的)
2. 如果遇到不一样的字符返回false
*/
func isPalindrome2(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !checkString(s[l]) {
			l++
		}
		for l < r && !checkString(s[r]) {
			r--
		}
		if l < r {
			if toLower(s[l]) != toLower(s[r]) {
				a, b, c, d := s[l], s[r], toLower(s[l]), toLower(s[r])
				fmt.Println(a, b, c, d)
				return false
			}
			l++
			r--
		}
	}
	return true
}

func checkString(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
