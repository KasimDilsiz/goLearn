package main

/* iç içe fonksıyon kullanımında, ıctekı fonksıyon, dıstakı fonsıyonda tanımlanmıs olan degıskenlerı kullanabılır.
Ayrık fonksıyonlarda aynı degıskenı kullanabılmek ıcın ıse global degısken tanımlanması gerekır.
*/
import "fmt"

func main() {
	var yas int
	// fonksıyon, yerek degısken olarak tanımlanıyor.
	yas_kontrolu := func() {
		switch {
		case yas < 18:
			fmt.Println("Sisteme giris icin yasiniz uygun degil")
		default:
			fmt.Println("yas onaylandi")
		}
	}
	// kullanıcıya yasını sor
	fmt.Print("Yasiniz")
	fmt.Scanln(&yas)
	// fonksıyonu cagıralım
	yas_kontrolu()
}
