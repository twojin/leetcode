package dp

var maxmin int = -1

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return 0
	}
	maxmin = amount + 1
	count := make([]int, amount)
	return dp(coins, amount, count)
}

func dp(coins []int, rem int, count []int) int {
	if rem < 0 {
		return -1
	}
	if rem == 0 {
		return 0
	}

	if count[rem-1] != 0 {
		return count[rem-1]
	}

	min := maxmin
	for i := 0; i < len(coins); i++ {
		res := dp(coins, rem-coins[i], count)
		if res >= 0 && res < min {
			min = res + 1
		}
	}

	if min == maxmin {
		count[rem-1] = -1
	} else {
		count[rem-1] = min
	}
	return count[rem-1]
}
