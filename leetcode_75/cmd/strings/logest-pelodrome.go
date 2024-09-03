package main

import "fmt"

func main() {

	// nums := []int{3, 2, 1, 5, 6, 4}
	// findKthLargest(nums, 2)
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cc"))

}

func longestPalindrome(s string) string {
	resStr := ""
	resLen := 0
	for i := range s {
		// odd
		findPalStr(s, i-1, i+1, &resLen, &resStr)
		//even
		findPalStr(s, i, i+1, &resLen, &resStr)

		// fmt.Println("s", string(s), i, resStr, resLen)

	}
	return resStr
}

func findPalStr(s string, l int, r int, maxLen *int, subStr *string) {
	for r < len(s) && l >= 0 && s[l] == s[r] {
		if *maxLen < r-l+1 {
			*maxLen = r - l + 1
			*subStr = s[l : r+1]
		}
		r++
		l--
	}
}
