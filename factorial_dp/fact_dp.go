package main

import "fmt"

func Fact(n int, dp []int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = n * Fact(n-1, dp)
	return dp[n]
}

func main() {
	var n = 5
	var dp = make([]int, n+1)
	for i := range n + 1 {
		dp[i] = -1
	}
	fmt.Println("Factorial of 5 using dp is",Fact(n,dp));
}