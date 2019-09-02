package main

// Kesitlerin eleman sayisi ve kapasitesini öğrenmek için; len ce cap fonksiyonlari kullanılır.
import "fmt"

func main() {
	dizi := [4]string{"a", "b", "c", "d"}
	kesit := dizi[:3]
	fmt.Printf("1) elaman sayisi=%d, kapatise=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "x")
	fmt.Printf("2) eleman sayisi=%d , kapatise=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = append(kesit, "y")
	fmt.Printf("3) eleman sayisi=%d , kapatise=%d, %v\n",
		len(kesit), cap(kesit), kesit)
	kesit = kesit[2:]
	fmt.Printf("4) eleman sayisi=%d , kapatise=%d, %v\n",
		len(kesit), cap(kesit), kesit)
}
