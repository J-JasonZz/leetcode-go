package main

import "fmt"

func generate(numRows int) [][]int {
	var res [][]int

	for i := 1; i <= numRows; i++ {
		list := make([]int, i)
		if i == 1 {
			list = []int{1}
		} else if i == 2 {
			list = []int{1, 1}
		} else {
			lastList := res[i - 2]
			for j := 0; j < i; j++ {
				if j == 0 || j == i - 1 {
					list[j] = 1
				} else {
					list[j] = lastList[j - 1] + lastList[j]
				}
			}
		}

		res = append(res, list)
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}