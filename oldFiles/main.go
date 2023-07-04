package main

import (
	"fmt"
	"strings"
)

var OrderNum int = 1234

func main() {

	option2 := "2018 Bergamont City"
	option1 := "2020 Gazelle Medeo"
	grouped := option1 + " " + option2
	splits := strings.Split(grouped, " ")
	fmt.Println(splits)

	var option2new string
	var option1new string

	for i := 0; i < len(option2); i=i+3 {
		if option1[i] != option2[i] {

			option2new += string(option2[i:])
			option1new += string(option1[i:])
			break
		}
	}
	fmt.Println(option2new)
	fmt.Println(option1new)
}
