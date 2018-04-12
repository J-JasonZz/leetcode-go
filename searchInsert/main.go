package main

import "fmt"

func main() {
	s := []int{1,3}
	fmt.Println(searchInsert(s, 3))
}

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}
	if target > nums[len(nums) - 1] {
		return len(nums)
	}
	if target == nums[len(nums) - 1] {
		return len(nums) - 1
	}

	for i, v := range nums {
		if target <= v {
			return i
		}
	}
	return 0
}