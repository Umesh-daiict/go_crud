package main

import (
	"fmt"
)

func main() {
	nums, k := []int{1, 2, 3, 4}, 5
	maxOperations(nums, k)
}

func maxOperations(nums []int, k int) int {
	left, right, pair := 0, len(nums)-1, 0
	cacheArr := make([]int, 0)
	for left < right {
		for left < right {
			if left != right && ((nums[left] + nums[right]) == k) {
				lCach := len(cacheArr) - 1
				fmt.Println("l1l_", lCach, cacheArr, nums[left], nums[right])

				if lCach == -1 || (lCach >= 0 && cacheArr[lCach] != nums[left] && cacheArr[lCach] != nums[right]) {
					pair++
					cacheArr = append(cacheArr, nums[right])
				}
				// fmt.Println(cacheArr, lCach)

			}
			left++
		}
		left = 0
		right--
	}
	return pair
}
