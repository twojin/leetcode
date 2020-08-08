package array

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/array-and-string/c5tv3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 根据二维数组的第一个元素排序
	quickSort(intervals)

	var merged [][]int
	for i := 0; i < len(intervals); i++ {
		if len(merged) == 0 || intervals[i][0] > merged[len(merged)-1][1] {
			merged = append(merged, intervals[i])
		} else {
			if intervals[i][1] > merged[len(merged)-1][1] {
				merged[len(merged)-1][1] = intervals[i][1]
			}
		}
	}

	return merged
}

// quickSort 快速排序 根据二维数组各元素中第一个排序
func quickSort(a [][]int) {
	l := len(a)
	if l <= 1 {
		return
	}

	midVal := a[0][0]
	left, right := 0, l-1
	for i := 1; i <= right; {
		if a[i][0] > midVal {
			a[i], a[right] = a[right], a[i]
			right--
		} else {
			a[i], a[left] = a[left], a[i]
			left++
			i++
		}
	}
	quickSort(a[:left])
	quickSort(a[left+1:])
}
