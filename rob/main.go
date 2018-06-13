package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 1, 4, 5}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}
	if size == 1 {
		 return nums[0]
	}
	a := nums[0]
	max := int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < size; i++ {
		temp := max
		max = int(math.Max(float64(a + nums[i]), float64(max)))
		a = temp
	}
	return max
}