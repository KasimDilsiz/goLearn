package main

import "fmt"

/*Struct ;programlama da birden fazla degiskenin bir arada paketlenmis hali gibi dusunebılırız.
yapı degıskenlerının alanlarına (özellıklerıne) erısmek ıcın nokta ısaretı kullanılır.or: oto1.fıyat=20000 gıbı
 vasıt ısmınde bir tür olusturalım . bu tur ıcınde 2 tamsayı 2 metın alanı olsun. */
type vasıta struct {
	yıl, fıyat int
	renk, tür  string
}

func main() {
	oto1 := vasıta{2010, 25000, "bordo", "4*4"}
	var tosbaa vasıta
	mx := vasıta{tür: "motorsiklet", yıl: 2010, fıyat: 15000}
	// yukarıda 3 tane yapı degıskenı olusturuldu
	tosbaa.yıl = 1980
	tosbaa.fıyat = 10000
	tosbaa.tür = "3kapı"
	//ekrana yazdıralım
	fmt.Println(oto1)
	fmt.Println(tosbaa)
	fmt.Println(mx)
	fmt.Println("toplam fiyat:", oto1.fıyat+tosbaa.fıyat+mx.fıyat)
}
