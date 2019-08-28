package main
//switch kullanımında birden fazla koşula uyan durum varsa ;sadece ilk uygun koşul işletilir.Diğerine bakılmaz.
import "fmt"

func main() {
	sayi:=0
	switch {
	case sayi>0:
		fmt.Println("sayimiz pozitif")
	case sayi<0:
fmt.Println("sayimiz negatiftir")
	default:
		fmt.Println("sayimiz 0'a eşittir")


	}
}
