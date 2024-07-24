package main

import "fmt"

type RecentCounter struct {
	previour []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	var result int
	this.previour = append(this.previour, t)
	for _, val := range this.previour {
		if val <= t && val >= t-3000 {
			result++
		}
	}
	return result
}

func main() {
	cur := Constructor()
	cur.Ping(1)
	cur.Ping(2)
	cur.Ping(3000)
	fmt.Println("c", cur.previour)
}
