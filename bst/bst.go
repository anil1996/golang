package main

import "fmt"

type node struct {
	key   int
	left  *node
	right *node
}

func (n *node) Insert(val int) {
	if val < n.key {
		if n.left == nil {
			n.left = &node{key: val}
		} else {
			n.left.Insert(val)
		}

	} else {
		if n.right == nil {
			n.right = &node{key: val}
		} else {
			n.right.Insert(val)
		}

	}
}
func (n *node) search(val int) bool {
	if n == nil {
		return false
	}
	if val < n.key {
		return n.left.search(val)
	} else if val > n.key {
		return n.right.search(val)
	}
	return true

}
func main() {
	node := &node{key: 100}
	node.Insert(200)
	node.Insert(90)
	node.Insert(190)
	node.Insert(900)
	fmt.Println(node)
	fmt.Println(node.search(90))
	fmt.Println(node.search(100))
	fmt.Println(node.search(9000))

}
