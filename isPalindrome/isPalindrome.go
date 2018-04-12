package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(isPalindrome(1001))
}

// 1234321   123321
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	i := x
	p := 1
	for i >= 10 {
		i = i / 10
		p = p * 10
	}

	j := x
	for p > 0 {
		head := j / p
		tail := j % 10
		if head != tail {
			return false
		}
		j = j % p / 10
		p = p / 100
	}

	math.Pow()
	return true
}