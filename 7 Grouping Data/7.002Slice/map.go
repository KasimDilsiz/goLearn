package main

import "fmt"

func main() {
	il_nufus := make(map[string]int)
	il_nufus["bitlis"] = 2144447
	il_nufus["istanbul"] = 78545633333
	il_nufus["izmir"] = 54444466823424
	il_nufus["yalova"] = 33354234
	fmt.Println(il_nufus["bitlis"])
	fmt.Println(il_nufus)
}

/* Dizilerde tutulan verilen saysal bir indeks ile erişilebilen verilerdi.harita ise sırasız ve indeksi olmayan
verilerle çalışmak için tasarlanmıstır.
ornek bı harıta ıse su sekılde tanımlanır .
var il_nufus map[string]int
il_nufus= make(map[string]int)
*/
