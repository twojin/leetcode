package array

import (
	"fmt"
	"testing"
)

func Test_removeDeplicates(t *testing.T) {

	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	len := removeDeplicates(nums)
	fmt.Println("len:", len)
	for i := 0; i < len; i++ {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
}
