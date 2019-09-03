package main

import "fmt"

func main() {
	panic("Ayyy")
	hata := recover() // bu satır calısmayacaktır.
	fmt.Println(hata)
}

/* panıc ıfadesının bulundugu satır ısletıldıgınde, fonksıyonun calısması kesılecektır.eger defer blogu olsaydı,
fonksıyon kesılmeden once orada yazanlar ısletılecektı.usttekı recover satırı defer blogu ıcerısınde olmadıgından,
asla ısletılmeyecektır.*/
