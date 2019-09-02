/*kesit, temel mantık olarak dizi yapısındadır.Diziden önemli farkı, boyutunun(eleman sayısının)
degistirilebilir olmasıdır.Kesitler,istenildıgınde boyutu kucultulen veya buyutulen dınamık dızılerdır.
kesitler dızıler gıbı tanımlanır. Ancak dizilerdekinin aksine boyut gırılmez.make fonksıyonu kesıt olsuturmak
icin kullanılır.
*/
package main

import "fmt"

func main() {
	slice := []int{1, 7, 0, 9}
	for i, deger := range slice {
		fmt.Printf("İndeks: %d Deger: %d\n ", i, deger)
	}
}
