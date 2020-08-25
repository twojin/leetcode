package array

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 1, 2}
	t.Log(majorityElement(nums))
	nums = []int{1, 2, 3, 1, 2, 2, 2, 1, 1, 1, 1, 1, 2, 2}
	t.Log(majorityElement(nums))
}
