package dp

import (
	"fmt"
)

func rob(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	s, d := 0, 0
	i := 0
	for i < len(nums)-1 {
		fmt.Println(i)
		s += nums[i]
		d += nums[i+1]
		i += 2
	}
	fmt.Println(i)
	if i < len(nums) {
		s += nums[len(nums)-1]
	}

	return max(s, d)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
