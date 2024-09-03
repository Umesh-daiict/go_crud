package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	queue := []*TreeNode{root}
	ans := []int{}
	if root == nil {
		return nil
	}
	for len(queue) > 0 {
		lq := len(queue)
		rightMost := &TreeNode{}
		for _, node := range queue {
			if node != nil {
				rightMost = node
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		queue = queue[lq:]

		if len(queue) > 0 {
			ans = append(ans, rightMost.Val)
		}
	}
	return ans
}
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	fmt.Println("maximum depth")
	fmt.Println("ans --", rightSideView(root))
}
