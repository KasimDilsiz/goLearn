package main

/* Fonksıyon tanımlanırken hem fonksıyonun alacagı degerler,hem de gerıye dondurulecegı degerler belirtilebilir.
bunun ıcın fonksıyonlar su sekılde tanımlanabılır:
func fonksıyon_ismi(gelen_deger int)(cevap int){...}
*/
import "fmt"

func main() {
	isim := isim_al()
	yas := yas_al()
	fmt.Printf("\nHosgeldiniz %s, Yasiniz: %d.", isim, yas)
	kontrol_et(yas)

}
func isim_al() (isim string) {
	fmt.Print("isminiz:")
	fmt.Scanln(&isim)
	return
}
func yas_al() (yas int) {
	fmt.Print("Yasiniz:")
	fmt.Scanln(&yas)
	return
}
func kontrol_et(yas int) {
	switch {
	case (yas < 18):
	}
	//if (yas < 18) {
	//	fmt.Println("Sisteme giris icin yasiniz müsait degil")
	//} else if (yas > 18) {
	//	fmt.Println("Sisteme basarili olarak giris yaptiniz");
	//	oturum_ac()
	//} else {
	//	fmt.Println("gırdıgınız aralik yanlistir")
	//}
}
func oturum_ac() {
	fmt.Println("isiniz bitince oturumu kapatmayi unutmayin")
}
