package main

import (
	"fmt"
)

func main() {
	fmt.Println(strStr("aaa", "a"))
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var start = 0
	for i, _ := range haystack {
		if haystack[i] == needle[0]  {
			if len(needle) == 1 {
				return start
			}
			start = i
			if i + len(needle) > len(haystack) {
				break
			}
			for j := 1; j < len(needle); j++ {
				if haystack[i + j] != needle[j] {
					break
				}
				if j == len(needle) - 1 {
					return start
				}
			}
		}
	}
	return -1
}