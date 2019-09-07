package main

import "fmt"

func main() {
	foo3()

	func() {
		fmt.Println("Anonymous func ran")
	}()
	func(x int) {
		fmt.Println("The meaning of life :", x)
	}(42)
}
func foo3() {
	fmt.Println("foo3 ran")
}
