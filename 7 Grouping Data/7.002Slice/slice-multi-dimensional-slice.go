package main

import "fmt"

func main() {
	jb := []string{"james", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)
	mb := []string{"Miss", "moneypenny", "strawberry", "Hazelnut"}
	fmt.Println(mb)

	xp := [][]string{jb, mb}
	fmt.Println(xp)
}
