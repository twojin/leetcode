package dp

func fib(N int) int {
	if N < 1 {
		return 0
	}
	memo := make(map[int]int)
	return helper(memo, N)
}

func helper(memo map[int]int, n int) int {
	if n <= 2 {
		return 1
	}

	if v, ok := memo[n]; ok {
		return v
	}

	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}
