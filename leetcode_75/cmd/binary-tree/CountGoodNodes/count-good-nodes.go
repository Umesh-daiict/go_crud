package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + dfsGood(root.Left, root.Val) + dfsGood(root.Right, root.Val)
}

func dfsGood(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}
	goods := 0
	if node.Val >= max {
		max = node.Val
		goods++
	}
	return goods + dfsGood(node.Left, max) + dfsGood(node.Right, max)

}
func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println("maximum depth")
	fmt.Println("ans --", goodNodes(root))
}
