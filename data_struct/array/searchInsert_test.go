package array

import (
	"testing"
)

func Test_searchIndex(t *testing.T) {
	nums1 := []int{1, 3, 5, 6}
	t.Log(searchInsert(nums1, 5))
	t.Log(searchInsert(nums1, 2))
	t.Log(searchInsert(nums1, 7))
	t.Log(searchInsert(nums1, 0))
}
