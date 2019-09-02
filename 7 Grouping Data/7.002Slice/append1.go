package main

import "fmt"

func main() {
	ilk := []int{1, 4}
	son := []int{5, 3}
	butun := append(ilk, son...)
	fmt.Printf("%v\n", butun)
}

/* Üç nokta işareti değişken sayıda argüman alabilen(variadic funcktions) fonksiyonlarda "tüm argümanlar,
tüm elemanlar" anlamında kullanılır.append(ilk,son...) satırının acıklaması,SON İSİMLİ KESİTİN TÜM ELEMANLARİNİ
İLK İSİMLİ KESİTE EKLE anlamındadır.
*/
