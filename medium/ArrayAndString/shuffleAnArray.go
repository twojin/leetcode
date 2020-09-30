package ArrayAndString

import (
	"math/rand"
)

type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	return Solution{data: nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.data
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	length := len(this.data)
	shuffle := make([]int, length)
	copy(shuffle, this.data)
	for i := 0; i < length; i++ {
		x := rand.Intn(i + 1)
		shuffle[i], shuffle[x] = shuffle[x], shuffle[i]
	}
	return shuffle
}
