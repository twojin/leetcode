package array

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	nums := []int{2, 1, 2, 1, 4}
	fmt.Println("single number:", singleNumber(nums))
}
