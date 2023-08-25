package part1

//给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
//
//单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。
//
//
//
//示例 1：
//
//输入：s = "Hello World"
//输出：5
//解释：最后一个单词是“World”，长度为5。
//示例 2：
//
//输入：s = "   fly me   to   the moon  "
//输出：4
//解释：最后一个单词是“moon”，长度为4。
//示例 3：
//
//输入：s = "luffy is still joyboy"
//输出：6
//解释：最后一个单词是长度为6的“joyboy”。

/*
*
1. 从后往前遍历字符串
2. 遇到非空格 用flag变量记录
3. 记录经过的空格数量
4. 字符长度减去空格数量以及遍历长度
*/
func lengthOfLastWord(s string) int {
	n := len(s)
	var flag bool
	spaceNum := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			spaceNum++
			if flag {
				return n - i - spaceNum
			}
		} else {
			flag = true
		}
	}
	return n - spaceNum
}
