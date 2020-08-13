package dp

func fib2(n int) int {
	if n < 1 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	prev, curr := 1, 1
	for i := 3; i <= n; i++ {
		sum := curr + prev
		prev = curr
		curr = sum
	}

	return curr
}
