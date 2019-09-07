package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := too(ii...)
	fmt.Println(n)
}
func too(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
