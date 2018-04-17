package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	nums := []int{-10,-3,0,5,9}
	fmt.Println(sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {
	return treeNode(nums, 0, len(nums) - 1)
}

func treeNode(nums []int, start int, end int) *TreeNode {

	if start > end {
		return nil
	}

	mid := (start + end) / 2
	var node = TreeNode{}
	node.Val = nums[mid]
	node.Left = treeNode(nums, start, mid - 1)
	node.Right = treeNode(nums, mid + 1, end)
	return &node
}