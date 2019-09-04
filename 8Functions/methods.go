package main

import "fmt"

type person struct {
	first string
	last  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
func main() {

	sa1 := secretAgent{
		person: person{
			"Kasim",
			"Dilsiz",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Mahmut",
			"Tuncer",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}
