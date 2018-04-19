package main

import "fmt"

func main() {
	prices := []int{1,2,3,4,5,6,1,5,1,6}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var sum int
	size := len(prices)
	for i := 0; i < size - 1; i++ {
		if prices[i + 1] > prices[i] {
			sum += prices[i + 1] - prices[i]
		}
	}
	return sum
}