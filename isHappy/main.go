package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(isHappy(10))
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 || n == 7 {
		return true
	}
	if n < 10 {
		return false
	}

	t := 0
	for n > 0 {
		t = t + int(math.Pow(float64(n % 10), 2.0))
		n = n / 10
	}
	return isHappy(t)
}