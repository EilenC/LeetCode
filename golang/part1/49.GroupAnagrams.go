package part1

/*
*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, s := range strs {
		t := []byte(s)
		sortBytes(t)
		m[string(t)] = append(m[string(t)], s) // sortedS 相同的字符串分到同一组
	}
	ans := make([][]string, 0, len(m)) // 预分配空间
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
func sortBytes(b []byte) {
	for i := 0; i < len(b); i++ {
		for j := i + 1; j < len(b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
}
