package main

import "fmt"

func fonk1(ülke *string) {
	// ülke="Türkiye"
	*ülke = *ülke + "\tCumhuriyeti"
}
func main() {
	metin := "Türkiye"
	fmt.Println(metin) // metin ="Türkiye
	fonk1(&metin)      // metin'in adresini fonksiyona bildir
	fmt.Println(metin) // metin = "Türkiye Cumhuriyeti
}
