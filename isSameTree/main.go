package main

import "fmt"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

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

     q3 := TreeNode{
          Val:3,
          Left:nil,
          Right:nil,
     }
     q2 := TreeNode{
          Val:2,
          Left:nil,
          Right:nil,
     }
     q1 := TreeNode{
          Val:1,
          Left:&q2,
          Right:&q3,
     }
     fmt.Println(isSameTree(&p1, &q1))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {

     if p == nil && q == nil {
          return true
     } else if p == nil || q == nil {
          return false
     }
     if p.Val == q.Val {
          return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
     }
     return false
}
