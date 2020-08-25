package array

func majorityElement(nums []int) int {
	var count, condition int
	for _, v := range nums {
		if count == 0 {
			condition = v
		}
		if condition == v {
			count++
		} else {
			count--
		}
	}
	return condition
}
