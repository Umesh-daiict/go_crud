package main

import (
	"fmt"
)

func dfs(curr int, visited []bool, isConnected [][]int) {
	visited[curr] = true
	for j := 0; j < len(isConnected); j++ {
		if isConnected[curr][j] == 0 {
			continue
		}
		if !visited[j] {
			dfs(j, visited, isConnected)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	count := 0
	visited := make([]bool, len(isConnected))
	for i := 0; i < len(isConnected); i++ {
		fmt.Println("visited ", i, visited)
		if !visited[i] {
			count++
			dfs(i, visited, isConnected)
		}
	}
	return count
}

func main() {
	fmt.Println("547. Number of Provinces")
	fmt.Println("0.", findCircleNum([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}))

}
