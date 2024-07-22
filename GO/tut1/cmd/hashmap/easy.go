package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Determine if Two Strings Are Close :", closeStrings("abc", "bca"), closeStrings("a", "aa"), closeStrings("cabbba", "abbccc"))

	// num1, num2 := []int{1, 2}, []int{1, 2, 2, 1, 1, 3}
	// fmt.Println("das", uniqueOccurrences(num1), uniqueOccurrences(num2))

	// num1, num2 := []int{1, 2, 3}, []int{2, 4, 6}
	// fmt.Println("ans", findDifference(num1, num2))
}

func closeStrings(word1 string, word2 string) bool {
	cmp1, cmp2 := make(map[rune]int), make(map[rune]int)
	for _, cm := range word1 {
		cmp1[cm]++
		if !strings.Contains(word2, string(cm)) {
			return false
		}
	}
	for _, cm := range word2 {
		cmp2[cm]++
	}
	fmt.Println("ds", cmp1, cmp2)
	valArr1, valArr2 := []int{}, []int{}
	for _, val := range cmp1 {
		valArr1 = append(valArr1, val)
	}

	for _, val := range cmp2 {
		valArr2 = append(valArr2, val)
	}
	slices.Sort(valArr1)
	slices.Sort(valArr2)
	return slices.Compare(valArr1, valArr2) == 0

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
