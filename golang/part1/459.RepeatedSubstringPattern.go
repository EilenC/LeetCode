package part1

import "strings"

//示例 1:
//
//输入: s = "abab"
//输出: true
//解释: 可由子串 "ab" 重复两次构成。
//示例 2:
//
//输入: s = "aba"
//输出: false
//示例 3:
//
//输入: s = "abcabcabcabc"
//输出: true
//解释: 可由子串 "abc" 重复四次构成。 (或子串 "abcabc" 重复两次构成。)

/*
*
1. 字符串长度奇数并且不全相等 直接返回false
2. 循环字符串长度一半 一直循环到长度 超过匹配字符串s或者完相等
*/
func repeatedSubstringPattern(s string) bool {
	ls := len(s)
	if ls%2 != 0 {
		var total int
		for _, v := range s {
			total += int(v)
		}
		if total%int(s[0]) != 0 {
			return false
		}
	}
	if ls == 2 {
		return s[0] == s[1]
	}
	var index int
	for {
		for i := 2 + index; i <= ls>>1; i++ {
			result := strings.Builder{}
			block := s[index:i]
			for j := 0; j < ls; j += i {
				result.WriteString(block)
			}
			if result.String() == s {
				return true
			}
		}
		if index >= ls {
			return false
		}
		index++
	}
}
