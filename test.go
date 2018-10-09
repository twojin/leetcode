package main

import (
	"fmt"
	"strconv"
)

var MinInt = -1 << 31
var MaxInt = 1<<31 - 1

func main() {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var i int
	for ; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if strs[j] == "" {
				return ""
			}

			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0][:i]
}

func countAndSay(n int) string {
	var say string = "1"
	if n <= 1 {
		return say
	}

	for n > 1 {
		var newStr string
		for i := 0; i < len(say); {
			count := 1
			num := string(say[i])
			j := i + 1
			for ; j < len(say) && say[i] == say[j]; j++ {
				count++
			}
			newStr = newStr + strconv.Itoa(count) + num
			i = j
		}
		say = newStr
		n--
	}

	return say
}
