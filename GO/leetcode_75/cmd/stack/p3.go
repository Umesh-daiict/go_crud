package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

type stack []int

func (s *stack) Push(n int) {
	*s = append(*s, n)
}
func (s *stack) Pop() int {
	lenn := len(*s) - 1
	res := (*s)[lenn]
	*s = (*s)[:lenn]
	return res
}

func (s *stack) Peek() int {

	return (*s)[len(*s)-1]

}

func main() {
	fmt.Println("394. Decode String")
	fmt.Println(
		decodeString("3[a]2[bc]"))
	decodeString("3[a2[c]]")
	decodeString("2[abc]3[cd]ef")

	// fmt.Println("735. Asteroid Collision")
	// fmt.Println("ans", asteroidCollision([]int{5, 10, -5}), asteroidCollision([]int{8, -8}), asteroidCollision([]int{10, 2, -5}),
	// 	asteroidCollision([]int{-2, -1, 1, 2}), asteroidCollision([]int{-2, -2, 1, -2}))

	// fmt.Println(". Removing Stars From a String")
	// fmt.Println(removeStars("leet**cod*e"), removeStars("erase*****"))

}

func decodeString(s string) string {
	result := ""
	k := 0
	mulStack := []int{}
	strStack := []string{}
	fmt.Println("mul ,str>>", strStack)
	i := 0
	for i < len(s) {
		fmt.Println("i", s[i])
		if unicode.IsDigit(rune(s[i])) {
			fmt.Println("digit", s[i])
			k = k*10 + (int(s[i]) - '0')
			// mulStack = append(mulStack, int(s[i])-'0')
			fmt.Println("case 1,", mulStack)
		} else if s[i] == '[' {
			strStack = append(strStack, result)
			mulStack = append(mulStack, k)
			result = ""
			k = 0
		} else if s[i] == ']' {
			ls, lm := len(strStack)-1, len(mulStack)-1
			if ls >= 0 && lm >= 0 {
				result = strStack[ls] + strings.Repeat(result, mulStack[lm])
				fmt.Println("result", result)
				strStack = strStack[:ls]
				mulStack = mulStack[:ls]
			}
		} else {
			result += string(s[i])
		}
		i++
	}
	fmt.Println("stacks", strStack, mulStack)
	return result
}

func asteroidCollision(asteroids []int) []int {
	sk := stack{asteroids[0]}
	for i := 1; i < len(asteroids); i++ {
		top := sk.Peek()
		if (top > 0 && asteroids[i] > 0) || (top < 0 && asteroids[i] < 0) {
			sk.Push(asteroids[i])
		} else {
			if top > 0 && asteroids[i] < 0 {

				if math.Abs(float64(top)) < math.Abs(float64(asteroids[i])) {
					for len(sk) >= 0 && top < asteroids[i] {
						sk.Pop()
						top = sk.Peek()
					}
					if len(sk) == 0 || asteroids[i] == top {
						sk.Push(asteroids[i])
					}
				} else if math.Abs(float64(top)) == math.Abs(float64(asteroids[i])) {
					sk.Pop()
				}

			} else {
				sk.Push(asteroids[i])
			}
		}
	}
	return sk
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
