package main

import "fmt"

/* Fonksıyonlar ana programdan farklı bır bellek alanında calıstırılırlar.
fonksıyon ıcınde tanımlanan bır degısken, sadece o fonksıyon
ıcınde gecerlıdır .
*/
func main() {
	a := 13
	b := 31
	sonuc := topla(a, b)
	fmt.Println("toplam(fonksiyon cevabı)=", sonuc)
	// topla fonksıyonu ıkı sayının toplamını dondurur
}

func topla(a, b int) int {
	return a + b
}
