package main

import "fmt"

func main() {
	sozluk := make(map[string]string)
	sozluk["ksm"] = "Kasim DİLSİZ"
	sozluk["GPU"] = "GNU is not unix"
	sozluk["WWW"] = "WORLD WIDE WEB"
	sozluk["IP"] = "Internet Protokol"
	delete(sozluk, "GNU")
	if karsılık, sonuc := sozluk["GNU"]; sonuc {
		fmt.Println(karsılık)
	} else {
		fmt.Println("Arana Sozcuk Bulunamadı")
	}
}
