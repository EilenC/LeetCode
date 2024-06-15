package part1

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。


示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false


提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

/*
*
1. 遍历整个字符串,使用slice代表为栈,如果都是左括号则push对应的右括号进入栈
2. 如果遇到右括号则判断当前栈顶是否是相符合即可
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var stack []byte
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else {
			sl := len(stack)
			if sl == 0 || stack[sl-1] != s[i] {
				return false
			}
			stack = stack[:sl-1]
		}
	}
	return len(stack) == 0
}
