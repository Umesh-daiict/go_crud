package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(". Removing Stars From a String")
	fmt.Println(removeStars("leet**cod*e"), removeStars("erase*****"))

}

func removeStars(s string) string {
	stack := []string{}
	for _, val := range s {
		if val == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, string(val))
		}
	}
	return strings.Join(stack, "")
}
