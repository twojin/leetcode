package ArrayAndString

/*
最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba"也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

/********************* 动态规划 **********************/
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	ispr := make([][]bool, n, n)
	for i := 0; i < n; i++ {
		ispr[i] = make([]bool, n)
	}

	max, lo, hi := 0, 0, 0
	for i := 0; i < n; i++ {
		j := i
		for j >= 0 {
			if s[i] == s[j] && (i-j < 2 || ispr[j+1][i-1]) {
				ispr[j][i] = true
				if max < i-j+1 {
					lo = j
					hi = i
					max = i - j + 1
				}
			}
			j--
		}
	}

	return s[lo : hi+1]
}

/********************* 暴力解法 **********************/
func BFlongestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	pd := ""
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if isPalindrome(s[i:j]) && len(s[i:j]) > len(pd) {
				pd = s[i:j]
			}
		}
	}

	return pd
}

func isPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}

	lo, hi := 0, n-1
	for lo < hi {
		if s[lo] != s[hi] {
			return false
		}

		lo++
		hi--
	}

	return true
}
