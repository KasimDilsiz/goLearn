package main

import "fmt"

type person2 struct {
	first string
	last  string
}
type secretAgent1 struct {
	person
	ltk bool
}

func (s secretAgent1) speak() {
	fmt.Println("I am", s.first, s.last, "-the secretgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "-the person speak")
}

type human interface {
	speak()
}

func bar1(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrr,", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrr,", h.(secretAgent1).first)
	}
	fmt.Println("I was passed into bar,", h)
}

type hotdog int

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
	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	bar1(sa1)
	bar1(sa2)
	bar1(p1)

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	fmt.Println(y)

}
