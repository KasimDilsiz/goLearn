package main

import "fmt"

func main() {
	foo()
	bar("Kasim")
	s1 := woo("Miss Fortune")
	fmt.Println(s1)

}
func foo() {
	fmt.Println("helle from foo")
}
func bar(s string) {
	fmt.Println("hello", s)
}
func woo(st string) string {
	return fmt.Sprint("hello from woo,", st)

}
