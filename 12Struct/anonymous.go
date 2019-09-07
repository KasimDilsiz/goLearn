package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Kasim",
		last:  "DİLSİZ",
		age:   23,
	}
	fmt.Println(p1)
}
