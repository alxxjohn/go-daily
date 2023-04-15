package main

import (
	"fmt"
	_ "time"
)

var OrderNum int = 1234

func main() {

	carCount := 21
	group := 10
	total := carCount / group
	remainder := carCount % 10
	fmt.Println(total)
	fmt.Println(remainder)

}


