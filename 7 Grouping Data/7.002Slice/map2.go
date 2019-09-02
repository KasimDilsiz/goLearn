package main

import "fmt"

func main() {
	il_nufus := map[string]int{
		"bitlis":   2174772352,
		"istanbul": 54563985636453,
		"izmir":    865327542,
		"yalova":   2423423434,
	}
	for anahtar, deger := range il_nufus {
		fmt.Printf("il:%s, Nufus: %d\n", anahtar, deger)
	}
}
