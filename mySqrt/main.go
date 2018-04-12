package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(16))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	last := 0.0
	res := 1.0

	for res != last {
		last = res;
		res = (res + float64(x) / res) / 2.0;
	}
	return int(res)
}

// (x / 2 + 1)
//func mySqrt(x int) int {
//	if x <= 1 {
//		return x
//	}
//
//	i := 0
//	j := x / 2 + 1
//
//	for {
//		mid := (i + j) / 2
//		sq := mid * mid
//		if sq == x {
//			 return mid
//		} else if sq < x {
//			i = mid + 1
//			nextSq := i * i
//			if nextSq == x {
//				return i
//			} else if nextSq > x {
//				return i - 1
//			}
//		} else {
//			j = mid - 1
//			nextSq := j * j
//			if nextSq <= x {
//				return j
//			}
// 		}
//	}
//
//	return 1
//}