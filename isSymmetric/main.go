package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	p3 := TreeNode{
		Val:2,
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
	fmt.Println(isSymmetric(&p1))
}


func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Right, root.Left)
}


func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
	}
	return false
}