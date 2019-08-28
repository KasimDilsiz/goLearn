package main

import "fmt"
import "reflect"

func main() {
	var a=5;  var b=4; var c="golang"
	fmt.Println(a+b)
	fmt.Println(reflect.TypeOf(a),reflect.TypeOf(b),reflect.TypeOf(c))

	isim := "Kas"
	isim= "Kasim"
	fmt.Println(isim)

}


/*yukarıda kullandıgımız refleck paketını degıskenın turunu ogrenmek ıcın dahıl ettık.prınt satırındakı reflect paketı
TypeOf fonksıyonu ıle degıskenın turunu ekrana yazdırıyoruz./
 Degısken tanımlaması bıraz daha kısaltılabılır;var sozcucu kaldırılabılır. Degıskenler bu sekılde tanımlanacaksa ;
degıskenın sadece ılk kullanuldıgı ksımda = yerıne := kullanılmalıdır. Sonrakı tum atamalarda yıne = kullanılmaya devam/

/*cok sayıda degısken tanımlama gerektıgı zaman kısa yol olarak soyle bı sey yapabılırız
 const(
	sakarya=54
	bıtlıs=13
	ıstanbul=34
	)  /





