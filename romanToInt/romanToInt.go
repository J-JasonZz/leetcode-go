package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("XX"))
}

var m = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if i == 0 || m[string([]rune(s)[i])] <= m[string([]rune(s)[i-1])] {
			res += m[string([]rune(s)[i])]
		} else {
			res += m[string([]rune(s)[i])] - 2*m[string([]rune(s)[i-1])]
		}
	}
	return res
}
