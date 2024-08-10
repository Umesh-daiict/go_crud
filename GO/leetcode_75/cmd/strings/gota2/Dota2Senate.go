package main

import "fmt"

func main() {

	// nums := []int{3, 2, 1, 5, 6, 4}
	// findKthLargest(nums, 2)

	// fmt.Println(predictPartyVictory("RD"))
	// fmt.Println(predictPartyVictory("RDD"))
	// fmt.Println(predictPartyVictory("DDR"))
	fmt.Println(predictPartyVictory("DDRRR"))

}
func predictPartyVictory(senate string) string {
	// result := 0
	d, r := []int{}, []int{}

	for i, val := range senate {
		if val == 'D' {
			d = append(d, i)
		} else {
			r = append(r, i)
		}
	}
	n := len(senate)
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r, r[0]+n)
		} else {
			d = append(d, d[0]+n)
		}
		d = d[1:]
		r = r[1:]

	}
	if len(r) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}

}
