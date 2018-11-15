package string

import (
	"strconv"
)

/*
 报数
报数序列是指一个整照其中的整数的顺序进数序列，按行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。
*/

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
