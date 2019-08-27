package main

import "fmt"

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

/* const (
	_=iota
 	kb:=1 << (iota*10)
	mb:=1 << (iota*10)
	gb:=1 << (iota*10)

)
func main () {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)func main(){
}
BU KODUMUZLA DA AYNI CIKTIYI ALABILIYORUZ .
*/

