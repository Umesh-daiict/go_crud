package main

import "fmt"

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	if root == nil {
		return 0
	}
	maxSum := root.Val
	maxlevel := 1
	curLevel := 1
	for len(queue) > 0 {
		lq := len(queue)
		curSum := 0
		for _, node := range queue {
			if node != nil {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				curSum += node.Val
			}
		}

		fmt.Println("max", maxlevel)
		if curSum > maxSum {
			maxSum = curSum
			maxlevel = curLevel
		}
		curLevel++
		queue = queue[lq:]

	}
	return maxlevel
}
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: -8}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println("1161. Maximum Level Sum of a Binary Tree")
	fmt.Println("ans --", maxLevelSum(root))
}
