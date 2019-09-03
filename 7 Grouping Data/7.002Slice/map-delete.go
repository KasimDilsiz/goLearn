package main

import "fmt"

func main() {
	m := map[string]int{
		"James":        32,
		"Miss Fortune": 27,
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)
	delete(m, "Ian Fleming")
	fmt.Println(m)

	fmt.Println(m["Miss Fortune"])
	fmt.Println(m["Ian Fleming"])

}

