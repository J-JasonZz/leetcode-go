package main

import "fmt"

func main() {
    l3 := ListNode{
        Val:4,
        Next:nil,
    }
    l2 := ListNode{
        Val:3,
        Next:&l3,
    }
    l1 := ListNode{
        Val:1,
        Next:&l2,
    }

    n3 := ListNode{
        Val:4,
        Next:nil,
    }
    n2 := ListNode{
        Val:2,
        Next:&n3,
    }
    n1 := ListNode{
        Val:1,
        Next:&n2,
    }
    l := mergeTwoLists(&l1, &n1)
    for l != nil {
        fmt.Println(l)
        l = l.Next
    }
}

type ListNode struct {
    Val int
    Next *ListNode
}
// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    var l3 *ListNode

    if l1.Val < l2.Val {
        l3 = l1
        l3.Next = mergeTwoLists(l1.Next, l2)
    } else {
        l3 = l2
        l3.Next = mergeTwoLists(l1, l2.Next)
    }

    return l3

}