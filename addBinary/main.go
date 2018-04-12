package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1111"
	b := "10011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	s := ""

	if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}

	// 进位
	flag := false
	for i := len(a) - 1; i >= 0; i-- {
		indexA := i
		indexB := i - (len(a) - len(b))
		charA, _ := strconv.Atoi(string([]rune(a)[indexA]))
		var charB int
		if indexB >= 0 {
			charB, _ = strconv.Atoi(string([]rune(b)[indexB]))
		} else {
			charB = 0
		}
		sum := charA + charB
		if flag {
			sum += 1
		}
		fmt.Println(sum)
		if sum == 1 || sum == 0 {
			s = strconv.Itoa(sum) + s
			flag = false
		}
		if sum == 2 {
			sum = 0
			s = strconv.Itoa(sum) + s
			flag = true
		}
		if sum == 3 {
			sum = 1
			s = strconv.Itoa(sum) + s
			flag = true
		}
		if i == 0 && flag {
			s = "1" + s
		}
	}
	return s
}
