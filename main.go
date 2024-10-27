package main

import (
	"fmt"
)

type Node struct {	
	value int
	left *Node
	right *Node
	middle *Node
}

func main() {
	fmt.Print("Hello\n")
	var tree = insert(nil, 100)
	tree = insert(tree, 100)
	tree = insert(tree, 0)
	tree = insert(tree, 200)	
	
	traverse(tree)
}

func traverse(t * Node) {
	if ( t == nil ) {
		return
	}
	fmt.Print("<\n")
	traverse(t.left)

	fmt.Print(t.value)

	for ( t.middle != nil) {
		t = t.middle
		fmt.Print(" -> ")
		fmt.Print(t.value)
	}

	fmt.Print(">\n")
	traverse(t.right)
}

func insert(t * Node, v int) *Node {
	if ( t == nil ) {
		return &Node{value:v, left:nil, middle:nil, right:nil}
	}

	if ( v < t.value ) {
		t.left = insert(t.left, v)
		return t
	}

	if ( v > t.value ) {
		t.right = insert(t.right, v)
		return t
	}

	if ( v == t.value) {
		var n = &Node{value:v, left:nil, middle:t.middle, right:nil}
		t.middle = n
	}

	return t
}

