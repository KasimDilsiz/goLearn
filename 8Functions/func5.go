package main

import "fmt"

func main() {
	var a float32 = 5
	var b float32 = 2
	fmt.Println("Bölüm= ", böl(a, b))
}
func böl(a, b float32) float32 {
	return a / b
}
