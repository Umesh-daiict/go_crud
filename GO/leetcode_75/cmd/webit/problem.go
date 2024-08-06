package main

func main() {

	// nums := []int{3, 2, 1, 5, 6, 4}
	// findKthLargest(nums, 2)

}

// selection sort/ heap sort/ merge sort,
// tern arrray into max heap

// Q1: K’th Smallest/Largest Element in Unsorted Array
// Q2: Write a function to print ZigZag order traversal of a binary tree
// Q3: Equilibrium index of an array
// Q4: Write optimized code to find the longest palindrome in a string. please look for optimized code.
// Q5. write a code to find a number in string. solve the problem using binary search.

// func findKthLargest(nums []int, k int) int {
// 	return quickSelect(0, len(nums)-1, nums, len(nums)-k)
// }

// func quickSelect(l int, r int, nums []int, k int) int {
// 	pivot, p := nums[r], l
// 	for i := l; i < r; i++ {
// 		if nums[i] <= pivot {
// 			nums[i], nums[p] = nums[p], nums[i]
// 			p++
// 		}
// 	}

// 	nums[p], nums[r] = nums[r], nums[p]

// 	if p > k {
// 		return quickSelect(l, p-1, nums, k)
// 	} else if p < k {
// 		return quickSelect(p+1, r, nums, k)
// 	} else {
// 		return nums[p]
// 	}
// }
