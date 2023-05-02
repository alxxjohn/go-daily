package main

import "fmt"

type Node struct {
	value int
	next *Node

}

type LinkedList struct {
	head *Node
	len  int
}

func main() {

	data := "hel"

	for  i,j := 0, 1; i < len(data); i,j = i+1, j+1 {
		fmt.Println(j,i)

	}

}