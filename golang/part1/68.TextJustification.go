package part1

import "strings"

/*
*
给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。

你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。

要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。

文本的最后一行应为左对齐，且单词之间不插入额外的空格。

注意:

单词是指由非空格字符组成的字符序列。
每个单词的长度大于 0，小于等于 maxWidth。
输入单词数组 words 至少包含一个单词。

示例 1:

输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
输出:
[

	"This    is    an",
	"example  of text",
	"justification.  "

]
示例 2:

输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
输出:
[

	"What   must   be",
	"acknowledgment  ",
	"shall be        "

]
解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",

	因为最后一行应为左对齐，而不是左右两端对齐。
	第二行同样为左对齐，这是因为这行只包含一个单词。

示例 3:

输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
输出:
[

	"Science  is  what we",
	"understand      well",
	"enough to explain to",
	"a  computer.  Art is",
	"everything  else  we",
	"do                  "

]

提示:

1 <= words.length <= 300
1 <= words[i].length <= 20
words[i] 由小写英文字母和符号组成
1 <= maxWidth <= 100
words[i].length <= maxWidth
*/

func fillWords(words []string, bg, ed, maxWidth int, lastLine bool) string {
	wordCount := ed - bg + 1
	spaceCount := maxWidth + 1 - wordCount // 除去每个单词尾部空格，+1 是最后一个单词的尾部空格的特殊处理
	for i := bg; i <= ed; i++ {
		spaceCount -= len(words[i]) // 除去所有单词的长度
	}

	spaceSuffix := 1 // 词尾空格
	var spaceAvg, spaceExtra int
	if wordCount == 1 {
		spaceAvg = 1
		spaceExtra = 0
	} else {
		spaceAvg = spaceCount / (wordCount - 1)   // 额外空格的平均值
		spaceExtra = spaceCount % (wordCount - 1) // 额外空格的余数
	}

	var ans strings.Builder
	for i := bg; i < ed; i++ {
		ans.WriteString(words[i]) // 填入单词
		if lastLine {
			// 特殊处理最后一行
			ans.WriteString(" ")
			continue
		}
		spaceToAdd := spaceSuffix + spaceAvg
		if i-bg < spaceExtra {
			spaceToAdd++
		}
		ans.WriteString(strings.Repeat(" ", spaceToAdd)) // 根据计算结果补上空格
	}
	ans.WriteString(words[ed])                               // 填入最后一个单词
	ans.WriteString(strings.Repeat(" ", maxWidth-ans.Len())) // 补上这一行最后的空格
	return ans.String()
}

func fullJustify(words []string, maxWidth int) []string {
	var (
		ans    []string
		cnt    int
		bg     int
		length = len(words)
	)
	for i, word := range words {
		cnt += len(word) + 1
		if i+1 == length || (cnt+len(words[i+1]) > maxWidth) {
			ans = append(ans, fillWords(words, bg, i, maxWidth, i+1 == length))
			bg = i + 1
			cnt = 0
		}
	}
	return ans
}
