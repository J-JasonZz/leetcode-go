package main

import "fmt"

func main() {
	p3 := TreeNode{
		Val:3,
		Left:nil,
		Right:nil,
	}
	p2 := TreeNode{
		Val:2,
		Left:nil,
		Right:nil,
	}
	p1 := TreeNode{
		Val:1,
		Left:&p2,
		Right:&p3,
	}
	fmt.Println(maxDepth(&p1))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	fmt.Println(leftDepth, rightDepth)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
