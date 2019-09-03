package main

// İsaretci degisken olusturmak ıcın, new fonksıyonu da kullanılabılır.
import "fmt"

func main() {
	adres := new(int)
	*adres = 2016
	fmt.Println("Adres:", adres)
	fmt.Println("Adres'teki Veri :", *adres)
}
