package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fp := head
	sp := head.Next
	fp.Next = nil
	for sp != nil {
		temp := sp.Next
		sp.Next = fp
		fp = sp
		sp = temp
	}
	head = fp
	return head
}
func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next = &ListNode{Val: 3}

	head.Next.Next.Next = &ListNode{Val: 4}
	cur := head

	for cur != nil {
		fmt.Println("val ", cur.Next)
		cur = cur.Next
	}

	head = reverseList(head)
	cur = head
	for cur != nil {
		fmt.Println("r val ", cur.Val)
		cur = cur.Next
	}
}
