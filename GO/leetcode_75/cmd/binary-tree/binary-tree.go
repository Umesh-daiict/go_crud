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
		fmt.Printf(`|| word: %s count : %d ||`, cur.word, cur.count)
	}
	if cur.left != nil {
		fmt.Printf("      \n / ")

		cur.left.printSubTree()
	}
	if cur.right != nil {
		fmt.Printf("\n       \\")
		cur.right.printSubTree()
	}

}

func (cur *tnode) printLevelTree() {
	if cur != nil {
		fmt.Printf(`|| word: %s count : %d ||`, cur.word, cur.count)
	}

	if cur.left != nil {
		fmt.Printf("      \n / ")

		cur.left.printSubTree()
	}
	if cur.right != nil {
		fmt.Printf("\n       \\")
		cur.right.printSubTree()
	}

}

type Binarytree struct {
	root *tnode
}

func (tr *Binarytree) insert(data string) {
	tr.root = tr.root.lUpdate(data)
}

func (tr *Binarytree) printLevelTree() {
	queue := []*tnode{}
	ans := [][]string{}
	// cur := tr.root
	if tr.root != nil {
		queue = append(queue, tr.root)
	}

	for len(queue) > 0 {
		lq := len(queue)
		level := []string{}
		// fmt.Printf(`|| word: %s count : %d ||`, tr.root.word, tr.root.count)
		for i := 0; i < lq; i++ {
			if queue[i] != nil {
				level = append(level, queue[i].word)
				queue = append(queue, queue[i].left)
				queue = append(queue, queue[i].right)
			}
		}
		if len(level) > 0 {
			ans = append(ans, level)
		}
		queue = queue[lq:]
	}
	fmt.Println("ans -->", ans)

}

func main() {
	tr := &Binarytree{}
	fileLines := readLines()
	for _, line := range fileLines {
		tr.insert(line)
	}
	tr.printLevelTree()

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
