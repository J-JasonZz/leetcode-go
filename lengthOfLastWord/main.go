package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("day"))
}

func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := string([]rune(s)[i])
		if c != " " {
			for j := i; j >= 0; j-- {
				a := string([]rune(s)[j])
				if a == " " {
					return count
				}
				count++
				if j == 0 {
					return count
				}
			}
		}
	}
	return count
}