package string

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	i := 0
	for ; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j] == "" {
				return ""
			}

			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:i]
}
