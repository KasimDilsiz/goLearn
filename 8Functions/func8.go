/*package main

// Tekrarlamalı fonksıyonlarda sonlandırma kosulu onemlıdır.Dikkat edılmezse,fonksıyon sonsuz donguye gırebılır.
import "fmt"

func main() {
	fmt.Println(faktöriyel(5))
}
func faktöriyel(sayi int) int {
	if sayi == 0 {
		return 1
	}
	return sayi * faktöriyel(sayi-1)
}
*/
// aynı fonksıyonun tekrarlamalı fonksıyon kullanılmadan yazılmıs halı.
package main

import "fmt"

func main() {
	fmt.Println(faktorıyel(4))
}
func faktorıyel(sayi int) int {
	carpım := 1
	for i := 1; i <= sayi; i++ {
		carpım = carpım * i
	}
	return carpım
}
