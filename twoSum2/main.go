package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	headIndex := 0
	footIndex := len(numbers) - 1
	for headIndex < footIndex {
		head := numbers[headIndex]
		foot := numbers[footIndex]
		sum := head + foot
		if sum == target {
			 return []int{headIndex + 1, footIndex + 1}
		} else if sum > target {
			footIndex--
		} else {
			headIndex++
		}
	}
	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	fmt.Println(twoSum(numbers,9))
}