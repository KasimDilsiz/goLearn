package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	sozluk := make(map[string]string)
	sozluk["ksm"] = "Kasim DİLSİZ"
	sozluk["GPU"] = "GNU is not unix"
	sozluk["WWW"] = "WORLD WIDE WEB"
	sozluk["IP"] = "Internet Protokol"
	tum_elemanlar, _ := json.MarshalIndent(sozluk, "", "")
	fmt.Println(string(tum_elemanlar))

}

/*
JSON ismindeki verı sunma biçimi,günümüzde farklı platformlar arasında veri aktarımı konusunda en populer yontemlerden
bırısıdır.
Harita ıcındekı tum verılerı donguye sokmak ıcın, diziler ve kesıtlerde de kullandıgımız range fonksıyonu kullanılabılır
*/
