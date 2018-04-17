package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		 return false
	} else {
		currentSum := root.Val
		if currentSum == sum && root.Right == nil && root.Left == nil {
			return true
		} else {
			return hasPathSum(root.Left, sum - currentSum) || hasPathSum(root.Right, sum - currentSum)
		}

	}

	return false
}


func main() {
	p4 := TreeNode{
		Val:5,
		Left:nil,
		Right:nil,
	}
	p3 := TreeNode{
		Val:3,
		Left:nil,
		Right:nil,
	}
	p2 := TreeNode{
		Val:2,
		Left:&p4,
		Right:nil,
	}
	p1 := TreeNode{
		Val:1,
		Left:&p2,
		Right:&p3,
	}
	fmt.Println(hasPathSum(&p1, 1))
}