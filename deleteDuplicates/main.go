package main

import "fmt"

func main() {
	l4 := ListNode{
		Val:2,
		Next:nil,
	}
	l3 := ListNode{
		Val:2,
		Next:&l4,
	}
	l2 := ListNode{
		Val:2,
		Next:&l3,
	}
	l1 := ListNode{
		Val:1,
		Next:&l2,
	}

	l := deleteDuplicates(&l1)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}



 type ListNode struct {
 	  Val int
	  Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var l *ListNode
	l = head

	val := head.Val

	for head.Next != nil {
		if head.Next.Val != val {
			head = head.Next
			val = head.Val
		} else {
			head.Next = head.Next.Next
		}
	}

	return l
}