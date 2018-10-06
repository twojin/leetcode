package string

/*
编写一个函数，其作用是将输入的字符串反转过来。

示例 1:

输入: "hello"
输出: "olleh"
示例 2:

输入: "A man, a plan, a canal: Panama"
输出: "amanaP :lanac a ,nalp a ,nam A"
*/

func reverseString(str string) string {
	sr := []rune(str)
	if len(sr) <= 1 {
		return str
	}
	head, tail := 0, len(sr)-1
	for head < tail {
		sr[head], sr[tail] = sr[tail], sr[head]
		head++
		tail--
	}

	return string(sr)
}
