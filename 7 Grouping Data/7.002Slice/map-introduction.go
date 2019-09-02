package main

import "fmt"

func main() {
	m := map[string]int{
		"james":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)
	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE PRINT", v)
	}
}
