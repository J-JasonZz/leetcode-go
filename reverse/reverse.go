package main

import (
	"math"
)

func main() {
	print(reverse(-1563847412))
}

func reverse(x int) int {
	const min_int32 = -2147483648
	const max_int32 = 2147483647
	i := int(math.Abs(float64(x)))
	j := 0
	for i > 0 {
		j = j * 10 + i % 10
		i = i / 10
	}
	if j > max_int32 || j < min_int32 {
		return 0
	}
	if x >= 0 {
		return j
	} else {
		return -j
	}
}