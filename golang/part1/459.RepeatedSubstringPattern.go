package part1

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
1. 遍历字符串,从字符串第二位开始截断拼接截断字符串,当i走到字符串一半的时候若两个字符串拼接起来 == 远字符串则 true
*/
func repeatedSubstringPattern(s string) bool {
	length := len(s)
	if length == 1 {
		return false
	}
	return s[length/2:]+s[:length/2] == s
}
