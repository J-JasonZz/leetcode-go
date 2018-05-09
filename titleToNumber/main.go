package main

import "fmt"

func titleToNumber(s string) int {
	byteArr := []byte(s)
	base := int('A' - 1)
	count := 0
	for _, v := range byteArr {
		count *= 26
		count += int(v) - base
	}
	return count
}

func main() {
	fmt.Println(titleToNumber("AAA"))
}
