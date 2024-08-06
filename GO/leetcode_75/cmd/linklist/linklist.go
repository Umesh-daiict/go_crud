package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	result := &ListNode{Val: 0, Next: head}
	fp, sp := head, result
	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}
	fmt.Println("middl", sp)
	sp.Next = sp.Next.Next
	return result.Next
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next.Next = &ListNode{Val: 4}
	cur := &head
	for cur != nil {
		fmt.Println("val ", cur.Val)
		cur = cur.Next
	}

	deleteMiddle(&head)

	cur = &head
	for cur != nil {
		fmt.Println("val ", cur.Val)
		cur = cur.Next
	}
}
