package array

import (
	"fmt"
	"testing"
)

func Test_intersect(t *testing.T) {
	nums1 := []int{2, 1}
	nums2 := []int{1, 2}

	fmt.Println(intersect(nums1, nums2))
}
