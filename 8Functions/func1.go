package main

import "fmt"

// Fonksiyonlar ıhtıyac halınde tekrar tekrar cagırılabılır.
func main() {
	a := 13
	b := 31
	topla(a, b)

}

// topla isimli fonksıyon asagıda tanımlanıyor.
func topla(m, n int) {
	fmt.Println("Sayilarin Toplami:", m+n)
}
