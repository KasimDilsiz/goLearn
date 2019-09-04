package main

import "fmt"

type kuslar struct {
	tür, renk, isim string
	yas             int
}

func main() {
	var papagan [3]kuslar
	papagan[0] = kuslar{yas: 1, renk: "beyaz", isim: "Sero", tür: "British"}
	papagan[1].renk = "Gri"
	papagan[1].yas = 2
	papagan[1].tür = "Safir"
	papagan[1].isim = "Duman"
	// ekrana yazdıralım
	for _, i := range papagan {
		fmt.Println(i)
	}
}
