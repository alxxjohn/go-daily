package main

import "fmt"


func main() {

	var a *int
	var b int
	var c int

	b = 23
	a = &b
	c = *a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}