package main

import "fmt"

func main() {
	var toplam int
	bilgisayar := map[string]int{
		"Dizüstü":  23,
		"Masaüstü": 123,
		"Tablet":   8,
		"Sunucu":   3,
	}
	for _, deger := range bilgisayar {
		toplam = toplam + deger
	}
	fmt.Printf("toplam %d bilgisayar var .\n", toplam)
}
