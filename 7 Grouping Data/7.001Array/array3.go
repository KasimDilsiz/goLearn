/* Dizi elemanlarina deger atamanin en kisa yolu; kğme parantezleri icinde bir  satirda yazmaktir.
orn: aday:=[5]int {81,63,23,14,21}
Not: Go dilinde standart olarak gelen range fonksiyonu;cok sayida veri üzerinde for'u döndürmek icin kullanilir.
*/
package main

import "fmt"

func main() {

	aday := [5]int{81, 63, 52, 14, 21}
	for i, puan := range aday {
		if puan < 50 {
			fmt.Println(i+1, ". aday basarisiz oldu", puan)
		}
	}
}
