package array

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
