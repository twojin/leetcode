package array

// removeDeplicates 从一个排序后的数组删除重复元素
func removeDeplicates(nums []int) int {
	i, j := 0, 0
	l := len(nums)

	for i = j + 1; i < l; i++ {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	return j + 1
}

// [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
// i=1,j=0; nums[i]=0,nums[j]=0; continue;
// i=2,j=0; nums[i]=1,nums[j]=0; j++=1; nums[j]=nums[i];
