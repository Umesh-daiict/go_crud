package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head.Next}
	odd, even := head, head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = dummy.Next
	return head
}
func main2() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	fmt.Println("oddEvenList :")

	fmt.Println(oddEvenList(head))
	cur := head
	for cur != nil {
		fmt.Println("val", cur.Val)
		cur = cur.Next
	}
}
