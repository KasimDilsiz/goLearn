package main

import "fmt"

/* herhangi bir degıskenın onune & ısaretı gelırse , o degıskenın adresini temsıl eder
 sımgesı ısaretcı degıskenlerının onune gelebılır.Isaretcının gosterdegı adresteki veriyi temsil eder.
pointers allow you to share a value stored in some memory location .use pointers when
-you want to change the data at a location
-you don't want to pass arround a lot of data */
import "fmt"

func kare(isaretci *int) {
	fmt.Println(*isaretci * *isaretci)

}
func main() {
	sayi := 5
	kare(&sayi)
}
