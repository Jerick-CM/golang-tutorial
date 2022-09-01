package main

import "fmt"

func main() {
	// var number uint16 = 288
	// fmt.Println("Hello World!", number)
	var out string = fmt.Sprintf("Number: %07d is cool", 45)
	fmt.Println(out)
}
