package main

import "fmt"

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}

	// newmp := make(map[string]int) new way to make map

	// fmt.Println(mp["apple"])
	// mp["apple"] = 600
	// fmt.Println(mp)

	val, ok := mp["apple"]

	fmt.Println(val, ok)
	fmt.Println(mp)

}
