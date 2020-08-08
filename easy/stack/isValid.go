package stack

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}

	parentheses := "([{)]}"

	stack := make([]string, 0, length/2)

	for i := 0; i < length; i++ {
		switch s[i] {
		case parentheses[0], parentheses[1], parentheses[2]:
			stack = append(stack, string(s[i]))
		case parentheses[3]:
			if pop(&stack) != string(parentheses[3-3]) {
				return false
			}
		case parentheses[4]:
			if pop(&stack) != string(parentheses[4-3]) {
				return false
			}
		case parentheses[5]:
			if pop(&stack) != string(parentheses[5-3]) {
				return false
			}
		default:
			return false
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}

func pop(s *[]string) string {
	if len(*s) == 0 {
		return ""
	}

	top := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return top
}
