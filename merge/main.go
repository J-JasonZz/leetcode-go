package main

import "fmt"

func main() {
	nums1 := []int{1,2,3,4,5,0}
	nums2 := []int{2}
	merge(nums1, 5, nums2,1)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {

	end := m + n - 1
	m = m - 1
	n = n - 1

	for m >= 0 && n >= 0 {
		if nums2[n] > nums1[m] {
			nums1[end] = nums2[n]
			end--
			n--
		} else {
			nums1[end] = nums1[m]
			end--
			m--
		}

	}

	for n >= 0 {
		nums1[end] = nums2[n]
		end--
		n--
	}
}