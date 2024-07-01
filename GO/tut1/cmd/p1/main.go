package main

import (
	"fmt"
)

func main() {
	nums, k := []int{1, 12, -5, -6, 50, 3}, 4
	findMaxAverage(nums, k)
}

func findMaxAverage(nums []int, k int) float64 {
	maxSum := sumOfk(nums, k, 0)

	for i := 1; i < len(nums)-k; i++ {
		curSum := maxSum - float64(nums[i]) + float64(nums[i+k])
		if curSum > maxSum {
			maxSum = curSum
		}
		fmt.Println("cur", curSum, maxSum)
	}
	return maxSum / float64(k)
}

func sumOfk(nums []int, k int, start int) float64 {
	sum, i := 0.0, start

	for i < (start + k) {
		sum += float64(nums[start])
		i++
	}
	return sum
}
