package main

/* defer ifadesı sız ıstedıgınız zaman yazıyorsunuz ama o mutlaka cıkarken yapıyor o işi.
defer ifadesi kullanıldıgı anda, ifadede kullanılan degıskenlerın degerı ne ise o degerler defer ıcın sabıtlenır.
yanı deger degısse bıle, fonksıyon sonunda ıfade calısırken, ilk bastakı degerle calısır.
bırden fazla defer ifadesi varsa, SON GIREN ILK CIKAR mantıgı ıle calısır.
*/
import "fmt"

func main() {
	islem_yap()
}
func islem_yap() {
	i := 0
	defer fmt.Println(i) // 2.sırada ısletılecek
	i++
	defer fmt.Println(i) // 1.sırada ısletılecek
	fmt.Println("Bu satır defer'lerden once isletilir")
}
