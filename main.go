package main

import "fmt"

type Person struct {
	name string
	age int
}



func main(){


	var alex Person
	alex = Person{name: "alex", age: 29}

	var p *Person
	p = &alex

	fmt.Println(p.name)
	fmt.Println(p.age)

}