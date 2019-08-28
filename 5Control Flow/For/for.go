package main

import "fmt"

/* genel for gösterimi for ifade1; ifade2;ifade3{ }
ifade1: Baslangıc kosul ıfadesıdır.Sadece döngu basladıgında işletilir.
ifade2:kontrol kosuludur.Döngünün her calısmasından once kontrol edilir.
ifade3:ilerleme kuralıdır.dögünün her çalışmasından sonra kontrol edilir .
*/
func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
