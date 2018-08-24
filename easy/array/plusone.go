/*
给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

package array

func plusOne(digits []int) []int {
	l := len(digits) - 1
	if l < 0 {
		return digits
	}

	if digits[l] < 9 {
		digits[l] = digits[l] + 1
		return digits
	}

	ret := make([]int, 0)
	tag := 0
	for i := l; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] = digits[i] + tag
			ret = digits
			break
		} else {
			digits[i] = 0
			tag = 1
		}

		if i == 0 && tag > 0 {
			digits[i] = 0
			ret = append([]int{1}, digits...)
			break
		}
	}

	return ret
}
