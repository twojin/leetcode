package array

import (
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	nums2 := []int{1, 2, 3}

	t.Log(pivotIndex(nums1))
	t.Log(pivotIndex(nums2))
}
