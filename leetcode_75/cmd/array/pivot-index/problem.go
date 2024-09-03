package main

func main() {

	nums := []int{1, 7, 3, 6, 5, 6}
	pivotIndex(nums)

}

func pivotIndex(nums []int) int {
	sumL, sumR := 0, 0

	for _, v := range nums {
		sumR += v
	}

	for i, v := range nums {
		sumR -= v

		if sumL == sumR {
			return i
		}
		sumL += v
	}

	return -1
}

// Q3: Equilibrium index of an array
