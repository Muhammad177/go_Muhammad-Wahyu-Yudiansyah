package main

import (
	"fmt"
	"math"
)

func frogJump(h []int) int {
	n := len(h)
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0 && i-j <= 2; j-- {
			dp[i] = min(dp[i], dp[j]+abs(h[i]-h[j]))
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(frogJump([]int{10, 30, 40, 20}))         // output: 30
	fmt.Println(frogJump([]int{30, 10, 60, 10, 60, 50})) // output: 40
}
