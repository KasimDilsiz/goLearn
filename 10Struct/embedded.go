package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Kasim",
			last:  "DİLSİZ",
			age:   22,
		},
		ltk: true,
	}
	p2 := person{
		first: "Alex",
		last:  "DE SOUZA ",
		age:   38,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(p2.first, p2.last, p2.age)

}
