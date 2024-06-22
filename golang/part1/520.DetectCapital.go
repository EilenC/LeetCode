package part1

/**
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。



示例 1：

输入：word = "USA"
输出：true
示例 2：

输入：word = "FlaG"
输出：false


提示：

1 <= word.length <= 100
word 由小写和大写英文字母组成
*/

/*
*
1. 遍历整个字符串,通过ascii来判断是否大小写
2. 设置两个大小写标志,每次判断大小写时候同事判断是否已经有其他情况出现
*/
func detectCapitalUse(word string) bool {
	var lower bool
	var upper bool
	length := len(word)
	if length > 1 {
		if word[0] >= 97 && word[0] <= 122 {
			lower = true
		}
	}
	for i := 1; i < length; i++ {
		if word[i] >= 65 && word[i] <= 90 {
			if lower {
				return false
			}
			upper = true
		}
		if word[i] >= 97 && word[i] <= 122 {
			if upper {
				return false
			}
			lower = true
		}
	}
	return true
}
