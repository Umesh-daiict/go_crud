package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
	curr := [][]int{entrance}
	ans := 0

	if maze == nil || entrance == nil {
		return -1
	}
	for len(curr) > 0 {
		next_curr := [][]int{}
		for _, val := range curr {
			if val[0] == 0 || val[0] == len(maze)-1 {
				if val[0] != entrance[0] || val[1] != entrance[1] {

					return ans
				}
			}

			if val[1] == 0 || val[1] == len(maze[0])-1 {
				if val[0] != entrance[0] || val[1] != entrance[1] {

					return ans
				}
			}

			if val[0]+1 < len(maze) && maze[val[0]+1][val[1]] != '+' {
				next_curr = append(next_curr, []int{val[0] + 1, val[1]})
				maze[val[0]+1][val[1]] = '+'
			}

			if val[0] > 0 && maze[val[0]-1][val[1]] != '+' {
				next_curr = append(next_curr, []int{val[0] - 1, val[1]})
				maze[val[0]-1][val[1]] = '+'
			}

			if val[1] < len(maze[0]) && len(maze[0]) > val[1]+1 && maze[val[0]][val[1]+1] != '+' {
				next_curr = append(next_curr, []int{val[0], val[1] + 1})
				maze[val[0]][val[1]+1] = '+'
			}

			if val[1] > 0 && maze[val[0]][val[1]-1] != '+' {
				next_curr = append(next_curr, []int{val[0], val[1] - 1})
				maze[val[0]][val[1]-1] = '+'
			}
		}
		curr = next_curr
		ans++
	}
	return -1

}

func main() {
	fmt.Println("1926. Nearest Exit from Entrance in Maze")
	fmt.Println("0.", nearestExit([][]byte{{'.'}, {'.'}}, []int{1, 0}))

	fmt.Println("1.", nearestExit([][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}}, []int{1, 2}))
	fmt.Println("2.", nearestExit([][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}, []int{1, 0}))

}
