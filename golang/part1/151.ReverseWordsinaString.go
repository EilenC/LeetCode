package part1

/*
*
给你一个字符串 s ，请你反转字符串中 单词 的顺序。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

示例 1：

输入：s = "the sky is blue"
输出："blue is sky the"
Explain
示例 2：

输入：s = "  hello world  "
输出："world hello"
解释：反转后的字符串中不能存在前导空格和尾随空格。
示例 3：

输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

提示：

1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词

进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。
*/

/*
*
1. 使用自带的split分割出所有空格区分的单词
2. 从末尾循环单词若不是空格则写入行字符串,并且补上一个空格
3. 去除最尾端多余的空格即可
*/
//func reverseWords(s string) string {
//	arr := strings.Split(s, " ")
//	length := len(arr)
//	ans := strings.Builder{}
//	for i := length - 1; i >= 0; i-- {
//		if arr[i] != "" {
//			ans.WriteString(arr[i])
//			ans.WriteString(" ")
//		}
//	}
//	return strings.TrimSuffix(ans.String(), " ")
//}

/*
*
1. 从尾巴开始遍历,遇到空格后开始计算单词使用一个[]byte接收
2. 使用一个hasSpace用来标记当前是否已经有一个空格了避免连续空格干扰
3. 当遇到下一个非空格就开始反向遍历单词 []byte 将正确的单词写入 strings.builder 中 若单词[]byte中是有单词的则最后补充一个空格
*/
//func reverseWords(s string) string {
//	var (
//		ans      = strings.Builder{}
//		hasSpace = false
//		word     = make([]byte, 0)
//	)
//
//	length := len(s)
//	for i := length - 1; i >= 0; i-- {
//		if s[i] == ' ' && hasSpace {
//			continue
//		}
//		if hasSpace {
//			for j := len(word) - 1; j >= 0; j-- {
//				if word[j] == ' ' {
//					continue
//				}
//				ans.WriteByte(word[j])
//			}
//			if len(word) > 0 { //这里再次判断是否有词,才选择是否添加空格
//				ans.WriteByte(' ')
//			}
//			hasSpace = false
//			word = make([]byte, 0)
//		}
//		if s[i] == ' ' {
//			hasSpace = true
//		}
//		if !hasSpace { //若没有空格并且不是空格的byte写入单词
//			word = append(word, s[i])
//		}
//	}
//	if len(word) > 0 { //由于字符串最起始位置不一定是空格,则有可能导致还有一个word没有处理
//		for j := len(word) - 1; j >= 0; j-- {
//			if word[j] == ' ' {
//				continue
//			}
//			ans.WriteByte(word[j])
//		}
//	}
//	return strings.TrimRight(ans.String(), " ")
//}

func reverseWords(s string) string {
	// 将字符串转换为字节切片以进行原地操作
	bytes := []byte(s)

	// 反转整个字符串
	reverse151(bytes, 0, len(bytes)-1)

	// 反转每个单词
	start := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverse151(bytes, start, i-1)
			start = i + 1
		}
	}
	// 反转最后一个单词（如果有的话）
	reverse151(bytes, start, len(bytes)-1)

	// 移除多余的空格
	return removeExtraSpaces(bytes)
}

// 反转字节切片的指定范围
func reverse151(bytes []byte, start, end int) {
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}

// 移除多余的空格
func removeExtraSpaces(bytes []byte) string {
	i, j := 0, 0
	for j < len(bytes) {
		if bytes[j] != ' ' || (i > 0 && bytes[i-1] != ' ') {
			bytes[i] = bytes[j]
			i++
		}
		j++
	}
	// 去除末尾的空格
	if i > 0 && bytes[i-1] == ' ' {
		i--
	}
	return string(bytes[:i])
}
