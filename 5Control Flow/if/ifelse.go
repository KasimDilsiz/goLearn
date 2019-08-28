package main

import "fmt"

func main() {
	var k1, k2, k3 int
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {

			k1++

		} else if i%3 == 0 {
			k2++

		}
		if i%4 == 0 {
			k3++

		}
	}
	fmt.Printf("2:%d\t,3:%d\t,4:%d",
		k1, k2, k3)
}
