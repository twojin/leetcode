package string

import (
	"fmt"
)

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/

func firstUniqChar(s string) int {
	n := len(s)
	if n == 1 {
		return 0
	}

	az := make([]int, 26, 26)

	for i := 0; i < n; i++ {
		az[s[i]-'a']++
		fmt.Println(s[i], 'a')
	}

	for i := 0; i < n; i++ {
		if az[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1

	// n := len(s)
	// if n == 1 {
	// 	return 0
	// }

	// for i := 0; i < n; i++ {
	// 	var count int
	// 	for j := 0; j < n; j++ {
	// 		if s[i] == s[j] {
	// 			count++
	// 		}
	// 	}
	// 	if count == 1 {
	// 		return i
	// 	}
	// 	count = 0
	// }

	// return -1
}
