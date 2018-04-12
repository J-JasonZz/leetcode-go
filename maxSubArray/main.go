package main

import "fmt"

func main() {
	s := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(s))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := 0
	thisSum := 0
	p := nums[0]
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num > p {
			p = num
		}
		thisSum += num
		if thisSum > maxSum {
			maxSum = thisSum
		} else if (thisSum < 0) {
			thisSum = 0
		}
	}
	if p < 0 {
		return p
	} else {
		return maxSum
	}
}