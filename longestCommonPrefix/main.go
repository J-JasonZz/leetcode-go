package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"flow","flow","flow"}
	fmt.Println(longestCommonPrefix(s))
}

// "asddd" "asddd" "as43123"
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	shortest := strs[0]
	for i := 1; i < len(strs); i++ {
		temp := strs[i]
		if len(temp) < len(shortest) {
			shortest = temp
		}
	}

	var s []string = make([]string, 0, len(shortest))
	for i := 0; i < len(shortest); i++ {
		temp := string(shortest[i])
		for j := 0; j < len(strs); j++ {
			str := strs[j]
			cur := string(str[i])
			if temp != cur {
				return strings.Join(s, "")
			}
		}
		s = append(s, temp)
	}
	return strings.Join(s, "")
}
