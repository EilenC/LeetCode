package part1

import "bytes"

//给你一个字符串 s ，将该字符串中的大写字母转换成相同的小写字母，返回新的字符串。
//
//
//
//示例 1：
//
//输入：s = "Hello"
//输出："hello"
//示例 2：
//
//输入：s = "here"
//输出："here"
//示例 3：
//
//输入：s = "LOVELY"
//输出："lovely"

/*
*
1. 遍历字符串s 判断字符的ascii是否在[65,90]中(大写A-Z)
2. 字符如果 > 90 则已是小写，无需操作
3. 字符如果 <= 90 则是大写字符,讲ascii + 32
*/
func toLowerCase(s string) string {
	var b = bytes.Buffer{}
	for i := range s {
		if s[i] >= 65 && s[i] <= 90 {
			b.WriteByte(s[i] + 32)
			continue
		}
		b.WriteByte(s[i])
	}
	return b.String()
}
