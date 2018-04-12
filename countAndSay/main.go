package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Println(countAndSay(6))
}

//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//6.	 312211
func countAndSay(n int) string {
	if n <= 0 {
		return "-1"
	}

	res := "1"
	for i := 1; i < n; i++ {
		var s []string
		index := 0
		for index < len(res) {
			val := string([]rune(res)[index])
			count := 0
			for index < len(res) && string([]rune(res)[index]) == val {
				index++
				count++
			}
			s = append(s, strconv.Itoa(count))
			s = append(s, val)
		}
		res = strings.Join(s, "")
	}
	return res
}