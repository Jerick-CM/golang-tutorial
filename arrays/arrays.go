package main

import "fmt"

func main() {
	// var arr [5]int
	// arr[0] = 100
	// arr[4] = 900

	// fmt.Println(arr)

	// arr := [3]int{3, 4, 5}
	// fmt.Println(arr)

	// arr := [3]int{3, 4, 5}
	// fmt.Println(len(arr))

	// arr := [3]int{3, 4, 5}
	// sum := 0

	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }

	// fmt.Println(sum)

	arr := [3]int{3, 4, 5}
	sum := 0
	arr2D := [2][3]int{{1, 2, 3}, {3, 3, 4}}

	fmt.Println(arr2D[0][2])

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

}
