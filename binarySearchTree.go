package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	root *Node
	count int
	name string
}

func (t *Tree) addNode(n Node) {
	if t.root == nil {
		t.root = &n
	} else {
		t.root.addChildNode(n)
	}
	t.count++
}


type Node struct {
	value int
	leftChild *Node
	rightChild *Node
}


func (parent *Node) addChildNode(newNode Node) {
	switch {
	case (newNode.value < parent.value) && (parent.leftChild == nil):
		parent.leftChild = &newNode
	case (newNode.value > parent.value) && (parent.rightChild == nil):
		parent.rightChild = &newNode
	case (newNode.value < parent.value) && (parent.leftChild != nil):
		parent.leftChild.addChildNode(newNode)
	case (newNode.value > parent.value) && (parent.rightChild != nil):
		parent.rightChild.addChildNode(newNode)
	}
}

func (start *Node) visit() {
	if start.leftChild != nil {
		start.leftChild.visit()
	}
	fmt.Printf("\n - %v", start.value)
	if start.rightChild != nil {
		start.rightChild.visit()
	}
}

func (start *Node) search(value int) {
	if start.value == value {
		fmt.Printf("\nfound: %v", start.value)
	} else if (value < start.value) && (start.leftChild != nil) {
		start.leftChild.search(value)
	} else if (value > start.value) && (start.rightChild != nil) {
		start.rightChild.search(value)
	}
}

func main() {
	// tree
	tree := new(Tree)
	*tree = Tree{nil, 0, "<my first tree"}
	(*tree).name = "<oh, it's my second tree>"

	// add nodes
	for i := 1; i <= 111; i++ {
		randValue := rand.Intn(100)
		tree.addNode(Node{randValue, nil, nil})
	}

	// test output
	fmt.Printf("\ntree: %v", *tree)
	fmt.Printf("\ntree count: %v", tree.count)
	fmt.Printf("\nroot: %v", tree.root.value)
	fmt.Printf("\nfirst left: %v", tree.root.leftChild.value)
	fmt.Printf("\nfirst right: %v", tree.root.rightChild.value)
	fmt.Printf("\nsecond right: %v", tree.root.rightChild.rightChild.value)

	// ordered output
	(*tree).root.visit()
	(*tree).root.search(91)
}