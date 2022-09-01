package main

import (
	"fmt"
)

func main() {

	age := 5
	if age >= 18 {
		fmt.Println("You can ride alone!")

	} else if age >= 14 {
		fmt.Println("You can ride with guardian")
	} else {
		fmt.Println("You cant ride yet")
	}

}
