package main

import "fmt"

/* Go dilinde ihtiyac olmayan bir degeri hic kullanmak istemiyorsak, o degiskenin
yerine alt cizgi karakterl (_) koyulmalidir */
func main() {

	aday := [5]int{81, 63, 52, 14, 21}
	basarili := 0
	for _, puan := range aday {
		if puan >= 50 {
			basarili += 1

		}
	}
	fmt.Println(" basarili aday sayisi", basarili)
}
