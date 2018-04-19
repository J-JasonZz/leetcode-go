package main

import "fmt"

func getRow(rowIndex int) []int {

	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	} else {
		row := getRow(rowIndex - 1)
		row = append(row, 1)
		size := len(row)
		start := size - 2
		end := size / 2
		if size % 2 == 0 {
			for i := start; i >= end; i-- {
				row[size - i - 1] = row[i] + row[i - 1]
			}
			for i := end ; i < size - 1; i++ {
				row[i] = row[size - i - 1]
			}
		} else {
			for i := end; i >= 1; i-- {
				row[i] = row[i] + row[i - 1]
			}

			for i := end + 1; i < size - 1; i++ {
				row[i] = row[size - i - 1]
			}
		}

		return row
	}
	return []int{}
}

func main() {
	fmt.Println(getRow(0))
}