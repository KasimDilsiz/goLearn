package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi)
	fmt.Println("the total is", x)
}
func sum(x []int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for the item indeks position", i,
			"we  are now adding", v, "to the totel which is now", v)
	}
	fmt.Println("the total is", sum)
	return sum
}
