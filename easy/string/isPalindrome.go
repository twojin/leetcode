package string

/*
 验证回文字符串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s2 := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		elem := s[i]
		if !isNumOrLet(elem) {
			continue
		}

		if elem >= 'A' && elem <= 'Z' {
			elem += 32
		}
		s2 = append(s2, elem)
	}

	lo, hi := 0, len(s2)-1
	for lo < hi {
		if s2[lo] != s2[hi] {
			return false
		}
		lo++
		hi--
	}

	return true
}

func isPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}

	lo, hi := 0, len(s)-1
	for lo < hi {
		head := s[lo]
		tail := s[hi]

		if !isNumOrLet(head) {
			lo++
		}
		if !isNumOrLet(tail) {
			hi--
		}
		if !isNumOrLet(head) || !isNumOrLet(tail) {
			continue
		}

		if head >= 'A' && head <= 'Z' {
			head += 32
		}
		if tail >= 'A' && tail <= 'Z' {
			tail += 32
		}

		if head != tail {
			return false
		}
		lo++
		hi--
	}

	return true
}

func isNumOrLet(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}
