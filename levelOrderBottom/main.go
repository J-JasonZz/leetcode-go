package main

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
	levelOrderBottom(&p1)
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var queue  []*TreeNode
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		var list []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				list = append(list, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		if list != nil {
			res = append([][]int{list},res...)
		}
	}

	return res
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}