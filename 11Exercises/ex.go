package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first: "Kasim",
		last:  "DİLSİZ",
		favFlavors: []string{"cıkolata",
			"sut",
			"elma",
		},
	}
	p2 := person{
		first: "James",
		last:  "BOND",
		favFlavors: []string{"armut",
			"ayran",
			"kek",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}

}
