package main

import "fmt"

type Node struct {
	val int;
	next *Node
}

func (n* Node) AddToEnd(v int) *Node {
	if n== nil {
		return &Node{val:v, next:nil}
	}
	curr:=n
	for curr.next!=nil {
		curr= curr.next
	}
	curr.next = &Node{}
	curr.next.val = v
	curr.next.next = nil
	return n
}

func (n* Node) PrintList() {
	curr:= n
	for curr!= nil {
		fmt.Println(curr.val)
		curr= curr.next
	}
}


func main() {
	var head *Node
	head = head.AddToEnd(10)
	head = head.AddToEnd(20)
	head = head.AddToEnd(30)
	head.PrintList()
	
}