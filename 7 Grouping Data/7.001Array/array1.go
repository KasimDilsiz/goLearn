package main

import "fmt"

/* 5 kısılık ogrencı grubunun not ortalaması
Dizinin tüm elemanlarına sira ile erismek icin for döngüsü kullanilabilir. */
func main() {
	var ogrenci [5] float32
	var toplam float32
	ogrenci[0] = 85
	ogrenci[1] = 80
	ogrenci[2] = 75
	ogrenci[3] = 85
	ogrenci[4] = 75
	for i := 0; i <= 4; i++ {

		toplam += ogrenci[i]

	}
	fmt.Println("toplam", toplam)
	fmt.Println("ortalama esittir", toplam/5)

}
