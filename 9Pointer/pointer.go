package main

// herhangi bir degıskenın onune & ısaretı gelırse , o degıskenın adresini temsıl eder
//* sımgesı ısaretcı degıskenlerının onune gelebılır.Isaretcının gosterdegı adresteki veriyi temsil eder.
import "fmt"

func kare(isaretci *int) {
	fmt.Println(*isaretci * *isaretci)

}
func main() {
	sayi := 5
	kare(&sayi)
}
