package main

import "fmt"

func majorityElement(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	numsMap := make(map[int]int)
	for _, v := range nums{
		count, ok := numsMap[v]
		if ok {
			numsMap[v] = count + 1
			if count + 1 > size / 2 {
				 return v
			}
		} else {
			numsMap[v] = 1
		}
	}
	return 0
}

func main() {
	nums := []int{2,2,1,1,1,2,2}
	fmt.Println(majorityElement(nums))
}