package main

import "fmt"

func main() {
	s := []int{1}
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	var temp = nums[0]
	var i = 1
	for {
		v := nums[i]
		if temp == v {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
		temp = v
		if i >= len(nums) {
			break
		}
	}

	return len(nums)
}