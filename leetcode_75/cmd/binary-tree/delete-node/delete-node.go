package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// find min from right sub tree
		cur := root.Right

		for cur.Left != nil {
			cur = cur.Left
		}
		root.Val = cur.Val
		root.Right = deleteNode(root.Right, cur.Val)
	}
	return root
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 6, Left: &TreeNode{Val: 7}}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 2}

	fmt.Println("delete node")
	fmt.Println("ans --", deleteNode(root, 3))
	fmt.Println("sda", root.Right, root.Right.Right, root.Right.Left)

}
