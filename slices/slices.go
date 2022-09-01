package main

import "fmt"

func main() {
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3] // colon inside array denotes slice
	// fmt.Println(s[1:])
	// fmt.Println(cap(s))
	// // fmt.Println(s[0:])

	// var a []int = []int{5, 6, 7, 8, 9}
	// b := append(a, 10)
	// fmt.Println(b)

	a := make([]int, 5)
	fmt.Println(a)
	fmt.Printf("%T", a)

}
