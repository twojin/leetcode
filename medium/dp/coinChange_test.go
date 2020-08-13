package dp

import "testing"

func Test_coinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	t.Log(coinChange(coins, amount))
}
