package main

import "fmt"

// func canVisitAllRooms(rooms [][]int) bool {
// 	visited := map[int]bool{0: true}
// 	checking := rooms[0]

// 	for len(checking) > 0 {
// 		newCheck := []int{}
// 		for _, val := range checking {
// 			_, ok := visited[val]
// 			if !ok {
// 				newCheck = append(newCheck, rooms[val]...)
// 				visited[val] = true
// 			}
// 		}
// 		checking = newCheck
// 	}
// 	fmt.Println("vi", len(visited), len(rooms))
// 	if len(visited) == len(rooms) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

func canVisitAllRooms(rooms [][]int) bool {

	// a map (so-called dictionary) to record visited rooms
	// Note: golang has no native set, so here we use map as alternative plan.
	visited := make(map[int]bool)

	// ------------------------------------------
	var dfs func(curRoom int) bool

	dfs = func(curRoom int) bool {

		if _, exist := visited[curRoom]; exist {

			// base case also known as stop condition
			return false
		}

		// mark current room as visited
		visited[curRoom] = true

		// general case:
		for _, nextRoom := range rooms[curRoom] {

			// Visit next room in DFS
			dfs(nextRoom)
		}
		return false
	}

	// ------------------------------------------

	// Launch DFS at room_#0
	dfs(0)

	// Return true if all rooms are visited
	return len(visited) == len(rooms)

}

func main() {
	fmt.Println("841. Keys and Rooms")
	fmt.Println("answer -- ", canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	fmt.Println("answer 2-- ", canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))

}
