package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum2(root *TreeNode, targetSum int) int {

	var dfs func(node *TreeNode, csum []int) int

	dfs = func(ndoe *TreeNode, csum []int) int {
		cur_add := 0
		if ndoe == nil {
			return 0
		}
		csum = append(csum, ndoe.Val)
		cur_summ := 0
		for i := len(csum) - 1; i >= 0; i-- {
			cur_summ += csum[i]
			if cur_summ == targetSum {
				cur_add += 1
				break
			}
		}
		return cur_add + dfs(ndoe.Left, csum) + dfs(ndoe.Right, csum)
	}

	return dfs(root, []int{})
}

func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, []int{})
}

func dfs(root *TreeNode, target int, currentPath []int) int {
	if root == nil {
		return 0
	}

	currentPath = append(currentPath, root.Val)
	pathCount, pathSum := 0, 0

	for i := len(currentPath) - 1; i >= 0; i-- {
		pathSum += currentPath[i]

		if pathSum == target {
			pathCount++
		}
	}

	pathCount += dfs(root.Left, target, currentPath)
	pathCount += dfs(root.Right, target, currentPath)

	currentPath = currentPath[:len(currentPath)-1]

	return pathCount
}
func main() {
	fmt.Println(" Path Sum III")
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: -3, Left: &TreeNode{Val: 11}}
	root.Right = &TreeNode{Val: 5, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 3}, Left: &TreeNode{Val: -2}}, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}}
	fmt.Println("answer -- ", pathSum(root, 8))
	fmt.Println("answer -- ", pathSum2(root, 8))

}
