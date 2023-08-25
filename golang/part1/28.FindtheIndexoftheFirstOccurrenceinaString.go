package part1

import "strings"

/*
*
1. 直接使用 needle 切割, 如果数量 == 1 则无匹配项目返回 -1
2. 如果 切割数量 > 1 则起始位置就是 第一个切割字符的长度
*/
func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	sl := strings.SplitN(haystack, needle, 2)
	if len(sl) > 1 {
		return len(sl[0])
	}
	return -1
}
