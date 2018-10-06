package string

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
    ****
*/

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	az := make([]int, 26, 26)
	/*进阶:
	如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？*/
	// 建立一个长度为len(s)的数组，此时空间复杂度为O(n) make([]rune, len(s), len(s))...
	for i := 0; i < len(s); i++ {
		az[s[i]-'a']++
		az[t[i]-'a']--
	}

	for i := 0; i < len(az); i++ {
		if az[i] != 0 {
			return false
		}
	}
	return true
}
