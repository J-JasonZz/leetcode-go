package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

func main() {
	//l4 := ListNode{
	//	Val:3,
	//	Next:nil,
	//}
	//l3 := ListNode{
	//	Val:3,
	//	Next:&l4,
	//}
	//l2 := ListNode{
	//	Val:1,
	//	Next:&l3,
	//}
	l1 := ListNode{
		Val:1,
		Next:nil,
	}
	fmt.Println(removeElements(&l1, 2))
}

func removeElements(head *ListNode, val int) *ListNode {
	var l *ListNode
	for head != nil {
		if l == nil && head.Val != val {
			l = head
			break
		} else {
			head = head.Next
		}
	}
	if l == nil || l.Next == nil {
		return l
	}

	head = l
	for head.Next.Next != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	if head.Next.Val == val {
		head.Next = nil
	}
	return l
}