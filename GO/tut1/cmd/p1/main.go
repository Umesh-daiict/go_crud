package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type lnode struct {
	line string
	prev *lnode
	next *lnode
}

type LinkList struct {
	head *lnode
	tail *lnode
}

func (ln *LinkList) insert(data string) {

	newNode := &lnode{line: data}
	if ln.head == nil {
		ln.head = newNode
		ln.tail = newNode
	} else {
		ln.tail.next = newNode
		newNode.prev = ln.tail
		ln.tail = newNode
	}
}
func (ln *LinkList) dispay() {
	cur := ln.head

	for cur != nil {
		fmt.Println(cur.line)
		cur = cur.next
	}
}

func (ln *LinkList) reverseDispay() {
	cur := ln.tail

	for cur != nil {
		fmt.Println(cur.line)
		cur = cur.prev
	}
}
func main() {
	var ln LinkList
	fileLines := readLines()
	for _, line := range fileLines {
		ln.insert(line)
	}
	ln.dispay()
	fmt.Println("reverse data")
	ln.reverseDispay()

}

func readLines() []string {
	lines := []string{}

	file, err := os.Open("../assets/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
