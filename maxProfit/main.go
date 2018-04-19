package main

import (
	"fmt"
	"math"
)

//func maxProfit(prices []int) int {
//	maxDiff := -1
//	size := len(prices)
//	for i := 0; i <= size - 2; i++ {
//		for j := i + 1; j <= size - 1; j++ {
//			currentDiff := prices[j] - prices[i]
//			if currentDiff > 0 && currentDiff > maxDiff {
//				 maxDiff = currentDiff
//			}
//		}
//	}
//	if maxDiff > 0 {
//		 return maxDiff
//	}
//	return 0
//}

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxDiff := 0
	for _, price := range prices {
		diff := price - minPrice
		if diff > maxDiff {
			 maxDiff = diff
		}
		if price < minPrice {
			 minPrice = price
		}
	}
	return maxDiff
}

func main() {
	prices := []int{5,2,4,5,6,1,3}
	fmt.Println(maxProfit(prices))
}