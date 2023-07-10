package main

import "fmt"



func main(){

	var names = [3] string {"hello", "tester", "other"}

	for _, elem  := range names {
		fmt.Println(elem)
	}

}