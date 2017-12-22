package main

import "fmt"

type node struct {
	root, left, right int
}

func main() {
	tree := node{1, 2, 3}
	fmt.Println(tree.print())
}

func (n node) print() (int, int, int) {
	return n.root, n.left, n.right
}
