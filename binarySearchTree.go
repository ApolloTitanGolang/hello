package main

import (
	"fmt"
	"math/rand"
	"encoding/json"
	"net/http"
	"time"
)

type Tree struct {
	Root  *Node	`json:"node"`
	Count int	`json:"count"`
	Name  string	`json:"name"`
}

func (t *Tree) addNode(n Node) {
	if t.Root == nil {
		t.Root = &n
	} else {
		t.Root.addChildNode(n)
	}
	t.Count++
}

type Node struct {
	Value int		`json:"value"`
	Count int		`json:"count"`
	LeftChild *Node		`json:"leftChild"`
	RightChild *Node	`json:"rightChild"`
}

func (parent *Node) addChildNode(newNode Node) {
	switch {
	case (newNode.Value < parent.Value) && (parent.LeftChild == nil):
		parent.LeftChild = &newNode
	case (newNode.Value > parent.Value) && (parent.RightChild == nil):
		parent.RightChild = &newNode
	case (newNode.Value < parent.Value) && (parent.LeftChild != nil):
		parent.LeftChild.addChildNode(newNode)
	case (newNode.Value > parent.Value) && (parent.RightChild != nil):
		parent.RightChild.addChildNode(newNode)
	}
}

func (start *Node) visit() {
	if start.LeftChild != nil {
		start.LeftChild.visit()
	}
	fmt.Printf("\n - %v - %v", start.Value, start.Count)
	if start.RightChild != nil {
		start.RightChild.visit()
	}
}

func (start *Node) search(value int) {
	if start.Value == value {
		fmt.Printf("\nfound: %v", start.Value)
	} else if (value < start.Value) && (start.LeftChild != nil) {
		start.LeftChild.search(value)
	} else if (value > start.Value) && (start.RightChild != nil) {
		start.RightChild.search(value)
	}
}

func getJSON(tree *Tree) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m, _ := json.Marshal(tree)
		fmt.Printf("\nJSON: %s", string(m))

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(m))
	}
}

func main() {

	// tree
	tree := new(Tree)
	*tree = Tree{nil, 0, "<my first tree"}
	(*tree).Name = "<oh, it's my second tree>"

	// add nodes
	for i := 1; i <= 20; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		randValue := rand.Intn(100)
		tree.addNode(Node{randValue, i,nil, nil})
	}

	// test output
	fmt.Printf("\ntree: %v", *tree)
	fmt.Printf("\ntree count: %v", tree.Count)
	fmt.Printf("\nroot: %v", tree.Root.Value)
	fmt.Printf("\nfirst left: %v", tree.Root.LeftChild.Value)
	fmt.Printf("\nfirst right: %v", tree.Root.RightChild.Value)
	fmt.Printf("\nsecond right: %v", tree.Root.RightChild.RightChild.Value)

	// ordered output
	(*tree).Root.visit()
	(*tree).Root.search(91)

	// start JSON Server
	http.HandleFunc("/", getJSON(tree))
	http.ListenAndServe(":8080", nil)
}