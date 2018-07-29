package array

import (
	"fmt"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	nums1 := []int{1, 2, 3, 1}
	nums2 := []int{1, 2, 3, 4}

	fmt.Println("nums1:", containsDuplicate(nums1))
	fmt.Println("nums2:", containsDuplicate(nums2))
}
