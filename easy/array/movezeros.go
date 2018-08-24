/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/

package array

func moveZeros(nums []int) {
	l := len(nums)
	if l <= 1 {
		return
	}

	for i := 0; i < l; i++ {
		ni := nums[i]
		if ni != 0 {
			continue
		}

		for j := i + 1; j < l; j++ {
			nj := nums[j]
			if nj == 0 {
				continue
			}
			nums[i] = nj
			nums[j] = 0
			break
		}
	}
}
