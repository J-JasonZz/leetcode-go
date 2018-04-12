package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValid("[()"))
}

func isValid(s string) bool {

	if len(s) == 0 {
		return false
	}
	f := "({["
	b := ")}]"

	var res []string
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if strings.Contains(f, c) {
			res = append(res, c)
		} else if strings.Contains(b, c) {
			if len(res) <= 0 {
				return false
			}
			last := string(res[len(res)-1])
			if strings.Index(f, last) != strings.Index(b, c) {
				return false
			} else {
				res = res[:len(res)-1]
			}
		}
	}
	if len(res) > 0 {
		return false
	}
	return true
}
