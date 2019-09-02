package main

/* Not: Dizilerdeki eleman sayisi , len(dizi_ismi) seklinde alinabilmektedir.
Dizi kullaniminda eleman sayisini len fonksiyonu ile bir degiskene atayip, hem
döngüde hem de diger islemlerde kullanmak isteyebiliriz .
*/import "fmt"

func main() {
	aday := [5]int{
		75,
		45,
		92,
		64,
		32,
	}
	var adaysayisi int = len(aday)
	for i := 0; i < adaysayisi; i++ {
		if aday[i] < 50 {
			fmt.Println(i+1, ". aday basarisiz oldu", aday[i])
		}

	}

}

