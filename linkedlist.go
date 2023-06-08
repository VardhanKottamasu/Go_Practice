package main

import (
	"fmt"
)

type node struct {
	data       int
	addr, head *node
}

func linkedList() {
	fmt.Print()
}

func insertNode(val int) {

	newNode := &node{data: val}
	if newNode.head == nil {
		newNode.head=newNode
	}

}
