package main

import "fmt"

func main() {
	i := 1
	toplam := 0
	/* Go'da c gibi dillerdeki gibi döngülerin aksine sadece tüm ihtiyaçlari karsılayacak sekılde bir for dongusu
	yapılmıstır.Şimdi for'un while gibi kullanımına bı örnek verelim . while döngüsünün mantığı, koşul sağlandığı sürece
	dön şeklindedir.
	*/

	for i <= 5 {
		i++
		toplam = toplam + i
	}
	fmt.Println("toplam:", toplam)
}
