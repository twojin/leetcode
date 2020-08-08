package array

import (
	"testing"
)

func Test_merge(t *testing.T) {
	intervals := [][]int{{15, 18}, {1, 3}, {8, 10}, {2, 6}}
	t.Log(merge(intervals))
}
