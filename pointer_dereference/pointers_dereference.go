package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed!"
}

func main() {
	/*
		x := 7
		y := &x
		fmt.Println(x, y)
		*y = 8 // derefrence and change the value of x
		fmt.Println(x, y)

	*/

	toChange := "hello"
	fmt.Println(toChange)
	// changeValue(&toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}
