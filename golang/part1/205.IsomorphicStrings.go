package part1

/*
*
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。

每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:

输入：s = "egg", t = "add"
输出：true
示例 2：

输入：s = "foo", t = "bar"
输出：false
示例 3：

输入：s = "paper", t = "title"
输出：true

提示：

1 <= s.length <= 5 * 104
t.length == s.length
s 和 t 由任意有效的 ASCII 字符组成
*/
func isIsomorphic(s string, t string) bool {
	sl := len(s)
	tl := len(t)
	if sl != tl {
		return false
	}

	m := make(map[byte]byte)
	for i := 0; i < sl; i++ {
		// 判断map存在是个存在key，表明是否已经对该字符建立过映射
		if tmp, ok := m[s[i]]; ok {
			//查看已经建立的映射的值和t对应位置的字符是否相同
			if tmp != t[i] {
				// 不相同说明又建立了一层对应关系
				// 即s中的一个字符与t中的两个字符建立了映射关系
				return false
			}
			continue
		}
		// 判断当前字符对应的t中的字符已经和s前面的某个字符建立过映射关系
		for _, v := range m {
			if v == t[i] {
				return false
			}
		}
		m[s[i]] = t[i]
	}

	return true
}
