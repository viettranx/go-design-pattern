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

type FollowerIterator interface {
	Next() Follower
	HasNext() bool
}

// sendMessage is used for any FollowerIterator
func sendMessage(iterator FollowerIterator, msg string) {
	for iterator.HasNext() {
		iterator.Next().Receive(msg)
	}
}

// Array Iterator

type FollowerArrayIterator struct {
	currentIdx int
	arr        []Follower
}

func NewFollowerArrayIterator(arr []Follower) FollowerIterator {
	return &FollowerArrayIterator{currentIdx: 0, arr: arr}
}

func (fi *FollowerArrayIterator) HasNext() bool {
	return len(fi.arr) > 0 && fi.currentIdx < len(fi.arr)
}

func (fi *FollowerArrayIterator) Next() Follower {
	flw := fi.arr[fi.currentIdx]
	fi.currentIdx++

	return flw
}

// Linked-List Iterator

type FollowerLinkedListIterator struct {
	node *LinkedNode
}

func NewFollowerLinkedListIterator(node *LinkedNode) FollowerIterator {
	return &FollowerLinkedListIterator{node: node}
}

func (fi *FollowerLinkedListIterator) HasNext() bool {
	return fi.node != nil
}

func (fi *FollowerLinkedListIterator) Next() Follower {
	node := fi.node
	fi.node = node.next

	return node.val
}

// Tree Iterator

type FollowerTreeStorage struct {
	node *TreeNode
}

func NewFollowerTreeStorage(node *TreeNode) FollowerTreeStorage {
	return FollowerTreeStorage{node: node}
}

func (flwTree FollowerTreeStorage) toArray(node *TreeNode) []Follower {
	if node == nil {
		return nil
	}

	followers := []Follower{node.val}

	for i := range node.children {
		followers = append(followers, flwTree.toArray(&node.children[i])...)
	}

	return followers
}

func (flwTree FollowerTreeStorage) toLinkedList(node *TreeNode, lNode *LinkedNode) *LinkedNode {
	if node == nil {
		return nil
	}

	lNode = &LinkedNode{val: node.val, next: lNode}

	for i := range node.children {
		lNode = flwTree.toLinkedList(&node.children[i], lNode)
	}

	return lNode
}

func (flwTree FollowerTreeStorage) Iterator() FollowerIterator {
	//return NewFollowerArrayIterator(flwTree.toArray(flwTree.node))
	return NewFollowerLinkedListIterator(flwTree.toLinkedList(flwTree.node, nil))
}

// Sample data & structures

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

func main() {
	message := "hello"

	fmt.Println("[a, b, b] Array iterator")
	iterator := NewFollowerArrayIterator(arrayOfFollowers)
	sendMessage(iterator, message)

	fmt.Println("a -> b -> c Linked-List iterator")
	iterator = NewFollowerLinkedListIterator(linkedListOfFollowers)
	sendMessage(iterator, message)

	fmt.Println("a -> [b -> [e, f], c]  Tree iterator")
	iterator = NewFollowerTreeStorage(treeOfFollowers).Iterator()
	sendMessage(iterator, message)
}
