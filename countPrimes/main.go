package main

import (
	"math"
	"fmt"
)

func main() {
	fmt.Println(countPrimes(6))
}

func countPrimes(n int) int {
	s := []int{}
	for i := 2; i < n; i++ {
		s = append(s, i)
	}
	limit := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= limit; i++ {
		if isPrimeNumber(i) {
			for j := 2; ; j++ {
				num := i * j
				if num >= n {
					break
				}
				s[num - 2] = 0
			}
		} else {
			continue
		}
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != 0 {
			count++
		}
	}
	return count
}

func isPrimeNumber(num int) bool {
	if num == 1 {
		 return false
	}
	if num == 2 {
		return true
	}
	limit := int(math.Ceil(math.Sqrt(float64(num))))
	for i := 2; i <= limit; i++ {
		for j := i; j < num; j++ {
			if i * j > num {
				break
			}
			if i * j == num {
				return false
			}
		}
	}
	return true
}