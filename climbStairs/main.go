package main

import "fmt"

func main() {

	fmt.Println(climbStairs(1))
}

func climbStairs(n int) int {

	res := make([]int, n + 2)
	res[0] = 0
	res[1] = 1
	res[2] = 2
	if n <= 2 {
		return res[n]
	}

	for i := 3; i <= n; i++ {
		res[i] = res[i - 2] + res[i - 1]	// 非原创
	}
	return res[n]
}