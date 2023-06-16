package main

import (
	"fmt"
)

type Follower interface {
	Receive(message string)
}

type Profile struct {
	name string
}

func (p Profile) Receive(message string) {
	fmt.Printf("%s has recieved message: %s\n", p.name, message)
}

var arrayOfFollowers = []Follower{
	Profile{name: "Peter"},
	Profile{name: "Mary"},
	Profile{name: "Tom"},
	Profile{name: "Henry"},
}

type LinkedNode struct {
	val  Follower
	next *LinkedNode
}

var linkedListOfFollowers = &LinkedNode{
	val: Profile{name: "Peter"},
	next: &LinkedNode{
		val: Profile{name: "Mary"},
		next: &LinkedNode{
			val:  Profile{name: "Tom"},
			next: nil,
		},
	},
}

type TreeNode struct {
	val      Follower
	children []TreeNode
}

var treeOfFollowers = &TreeNode{
	val: Profile{name: "Peter"},
	children: []TreeNode{
		{
			val: Profile{name: "Tom"},
			children: []TreeNode{
				{val: Profile{name: "Mary"}},
				{val: Profile{name: "Vincent"}},
				{val: Profile{name: "Vicky"}},
			},
		},
		{
			val: Profile{name: "Bob"},
			children: []TreeNode{
				{val: Profile{name: "Alice"}},
			},
		},
	},
}

func sendMessageForArray(msg string) {
	for i := range arrayOfFollowers {
		arrayOfFollowers[i].Receive(msg)
	}
}

func sendMessageForLinkedList(msg string) {
	node := linkedListOfFollowers

	for node != nil {
		node.val.Receive(msg)
		node = node.next
	}
}

func sendMessageForTree(node *TreeNode, msg string) {
	if node == nil {
		return
	}

	node.val.Receive(msg)

	for i := range node.children {
		sendMessageForTree(&node.children[i], msg)
	}
}

func main() {
	message := "hello"

	fmt.Println("Sending for array")
	sendMessageForArray(message)

	fmt.Println("Sending for linked-list")
	sendMessageForLinkedList(message)

	fmt.Println("Sending for tree")
	sendMessageForTree(treeOfFollowers, message)
}
