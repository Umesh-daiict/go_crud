package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type tnode struct {
	word  string
	count int
	left  *tnode
	right *tnode
}

func (np *tnode) lUpdate(data string) *tnode {
	newNode := &tnode{word: data, count: 1}
	if data == "i" {
		fmt.Println("Hii")
	}
	if np == nil {
		np = newNode
	} else {
		result := strings.Compare(np.word, data)
		if result == 0 {
			np.count++
		} else if result == 1 {
			if np.right == nil {
				np.right = newNode
			} else {
				np.right = np.right.lUpdate(data)
			}
		} else if result == -1 {
			if np.left == nil {
				np.left = newNode
			} else {
				np.left = np.left.lUpdate(data)
			}
		}
	}
	return np
}

func (cur *tnode) printSubTree() {
	if cur != nil {
		fmt.Println("word:", cur.word, " count :", cur.count)
	}
	if cur.left != nil {
		cur.left.printSubTree()
	}
	if cur.right != nil {
		cur.right.printSubTree()
	}

}

type Binarytree struct {
	root *tnode
}

func (tr *Binarytree) insert(data string) {
	tr.root = tr.root.lUpdate(data)
}

func main() {
	tr := &Binarytree{}
	fileLines := readLines()
	for _, line := range fileLines {
		tr.insert(line)
	}
	tr.root.printSubTree()

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
