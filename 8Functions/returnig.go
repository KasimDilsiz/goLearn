package main

import "fmt"

func main() {
	s2 := ab()
	fmt.Println(s2)
	x := ab()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)
}


func ab() func() int {
	return func() int {
		return 451
	}
}
