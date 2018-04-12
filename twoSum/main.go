package main

import "fmt"

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for index, num := range nums {
		m[num] = index
	}
	
	for _, num := range nums {
		v, ok := m[target - num]
		if ok {
			firstIndex := index(nums, num)
			secondIndex := v
			if firstIndex != secondIndex {
				return []int{firstIndex,secondIndex}
			}
		}
	} 
	
	return nums

}

func index(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return 0
}

func main() {
	nums := []int{1,2,3,4}
	fmt.Println(twoSum(nums, 7))
}