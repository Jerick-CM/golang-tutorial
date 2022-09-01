package main

import "fmt"

func main() {

	ans := 1

	switch ans {
	case 1:
		fmt.Println("1. one")
	case 2:
		fmt.Println("2. two")
	case 3:
		fmt.Println("3. three")
	case 4:
		fmt.Println("4. four")
	default:
		fmt.Println("not a case")
	}

	new_ans := -1

	switch {
	case new_ans > 0:
		fmt.Println("Greater than zero")
	case new_ans < 0:
		fmt.Println("Less than zero")
	default:
		fmt.Println("Zero")
	}

}
