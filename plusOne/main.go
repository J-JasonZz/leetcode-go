package main

import "fmt"

func main() {
	s := []int{9, 9, 9}
	fmt.Println(plusOne(s))
}

func plusOne(digits []int) []int {
	//var b := false
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]
		digits[i] = num + 1;
		if digits[i] < 10 {
			break;
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}