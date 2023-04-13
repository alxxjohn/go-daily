package test

import (
	"fmt"
	"time"
)

var OrderNum int = 1234

func main() {

	fmt.Println("Start")

	var runnerFrame string = "Testing"
	nameOfUser := "Alex"
	testt := sayHello("alex")

	fmt.Println(runnerFrame)
	time.Sleep(4)
	fmt.Print("next. \n")
	fmt.Println(OrderNum)
	fmt.Println(testt)
	fmt.Println(nameOfUser)

}

func sayHello(name string) string {
	return fmt.Sprintf("Hello %v", name)

}
