package main

import "fmt"

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	arr := []int{}
	cur := head
	var max int
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	fmt.Println("print --", arr)
	n := len(arr)
	fmt.Println("asd,", n/2-1)
	for i := 0; i <= (n/2 - 1); i++ {
		fmt.Println("index", i)
		twinsum := arr[i] + arr[n-1-i]
		if twinsum > max {
			max = twinsum
		}
	}
	return max
}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 1000}
	// head.Next.Next = &ListNode{Val: 3}
	// head.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next = &ListNode{Val: 3}

	// head.Next.Next.Next = &ListNode{Val: 4}
	cur := &head
	for cur != nil {
		fmt.Println("val ", cur.Val)
		cur = cur.Next
	}

	fmt.Println("answer --", pairSum(&head))

}
