package main

import (
	"fmt"
	"slices"
)

func main() {

	num1, num2 := []int{1, 2}, []int{1, 2, 2, 1, 1, 3}
	fmt.Println("das", uniqueOccurrences(num1), uniqueOccurrences(num2))
	// num1, num2 := []int{1, 2, 3}, []int{2, 4, 6}
	// fmt.Println("ans", findDifference(num1, num2))
}

func uniqueOccurrences(arr []int) bool {
	hashmap := make(map[int]int)
	dupmap := make(map[int]int)

	for _, val := range arr {
		hashmap[val]++
	}
	for _, val := range hashmap {
		dupmap[val]++
		if dupmap[val] > 1 {
			return false
		}
	}
	return true
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	res1, res2 := []int{}, []int{}
	hashmap := make(map[int]int)

	for _, val := range nums1 {
		hashmap[val] = 0
	}

	for _, val := range nums2 {
		_, exists := hashmap[val]

		if exists {
			hashmap[val] = 1
		} else {
			res2 = append(res2, val)
		}
	}

	for key, value := range hashmap {
		if value == 0 {
			res1 = append(res1, key)
		}
	}
	slices.Sort(res1)
	slices.Sort(res2)

	return [][]int{res1, res2}

}
