package main

import "fmt"

func kare(x int) {
	defer func() {
		hata := recover()
		if hata != nil { //recover metni var mi?
			fmt.Println(hata)
			fmt.Println("Panik Olustu. Defer kismi calisti")
		}
	}()
	panic("----durma noktasi----")
	fmt.Println(x * x)

}
func main() {
	kare(5)
}
