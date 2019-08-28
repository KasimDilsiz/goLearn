package main

import "fmt"

var y string

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSING a VALUE of that type to the variable
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)
	y = "Bond, James Bond"

	fmt.Println(y)
	fmt.Printf("%T", y)

}
