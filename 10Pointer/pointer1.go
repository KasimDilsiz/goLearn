package main

/* Fonksıyona ısaretcı gonderıldıgınde; degıskenın kopyası degıl,adresı gonderılır.
Fonksıyonların bu ısaretcıyı kullanarak yaptıgı her degısıklık, orjınal verı uzerınde yapılır*/
import "fmt"

func main() {
	sayi := 54
	fmt.Println("Sayı:", sayi)
	fmt.Println("Sayının Adresi:", &sayi)
	adres := &sayi
	fmt.Println("Adres:", adres)
	fmt.Println("Adresteki veri:", *adres)
}
