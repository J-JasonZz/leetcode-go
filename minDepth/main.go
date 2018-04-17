package main


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil{
		return 1
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	}
	minLeft := minDepth(root.Left) + 1
	minRight := minDepth(root.Right) + 1
	if minLeft < minRight {
		return minLeft
	} else {
		return minRight
	}
}

func main() {

}