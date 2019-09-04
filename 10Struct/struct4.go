package main

/* Metot'un -sekıl acısından- fonsksıyondan tek farkı;func sozgucu ile metodun ismi arasında parantez ıcınde
yazılan kısımdır.Bu kısım, fonksıyonu bır yapıya baglayarak metot halıne getırır.*/
import "fmt"

type insan struct {
	boy float32
}

func (birisi insan) ideal_kilo() float32 {
	kilo := 0.9*birisi.boy - 85
	return kilo
}
func main() {
	murat := insan{boy: 172}
	fmt.Println(murat.ideal_kilo())
}

/* func(birisi insan)ideal_kilo() float32 {
birisi : metot ıcınde gecıcı olarak kullanılacak olan ınsan turunde bır yerel degısken
insan :ALICI: bu metot ınsan yapısına baglı.
ıdeal_kilo: Metodun ısmı
float32: metot, cevap olarak ondalıklı sayı dondurecek */
