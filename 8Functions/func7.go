package main

/* Argüman sayısı değişen fonksıyonlar tanımlarken, alabılecegı degerlerın belirtildigi yerde üç nokta
isareti kullanılır.Or: func topla(argumanlar ...int)int{
*/
import "fmt"

func main() {
	fmt.Println(topla(-5, 4))
	fmt.Println(topla(1, 9, 2, 3))
}
func topla(argumanlar ...int) int {
	var toplam int
	for _, i := range argumanlar {
		toplam += i
	}
	return toplam
	// Fonksıyona gırecek arguman sayısını bılmedıgımız ıcın for i:=0;i<=5;i++ gibi bir yazım kullanamayız.
}
