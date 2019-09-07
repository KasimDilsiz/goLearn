package main

import "fmt"

func fonk(x *int) int {
	*x++
	return *x
}
func main() {
	sayi := 0
	fmt.Println(fonk(&sayi))
}
