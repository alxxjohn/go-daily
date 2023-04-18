package main

import (
	"fmt"
	"strings"
)

var OrderNum int = 1234

func main() {

	option2:= "Ford Pinto"
	option1 := "Ford Focus"
	grouped := option1 +  " "  + option2
	splits := strings.Split(grouped, " ")
	fmt.Println(splits)

	for i := 0; i < len(option2); i++ {

		fmt.Printf("%c\n", option1[i])

	}




}


