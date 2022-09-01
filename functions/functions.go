package main

import "fmt"

func test(x int, y int) {
	fmt.Println("test")
	fmt.Println(x + y)
}
func test2(x, y int) {
	// fmt.Println("test")
	fmt.Println(x + y)
}

func compute(x, y int) (int, int) {
	// fmt.Println("test")
	// fmt.Println(x + y)
	return x + y, x - y
}

func execute(x, y int) (z1 int, z2 int) {
	z1 = x + y
	z2 = x - y
	return
}

func trigger(x, y int) (z1 int, z2 int) {
	defer fmt.Println("Hello")
	z1 = x + y
	z2 = x - y
	return
}
func main() {

	// test(1, 11)
	// test2(4, 3)
	// ans1, ans2 := compute(5, 7)
	// a1, a2 := execute(5, 7)
	b1, b2 := trigger(5, 7)

	// fmt.Println("Sum is ", ans)
	// fmt.Println(ans1, ans2)
	// fmt.Println(a1, a2)
	fmt.Println(b1, b2)
	// fmt.Println(ans)

}
