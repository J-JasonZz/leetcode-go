package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3,2,2,3}
	fmt.Println(removeElement(s, 3))
}

func removeElement(nums []int, val int) int {
	sort.Ints(nums)
	var start int
	var end int
	count := 0
	for i, v := range nums {
		if v == val {
			if count == 0 {
				start = i
			}
			count++
			if v != val {
				break
			}
		}
	}
	end = start + count
	nums = append(nums[:start], nums[end:]...)
	return len(nums)
}