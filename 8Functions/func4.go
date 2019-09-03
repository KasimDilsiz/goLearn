package main

/*Hem ana fonksıyonda hem de alt fonksıyonlarda aynı degıskenı kullanmak
ıcın, değişken global olarak tanımlanmalıdır.*/
import "fmt"

// Global degıskenler fonksıyon dısında tanımlanır.
var a, b int = 1, 2

func degstir() {
	gecici := a
	a = b
	b = gecici
}
func main() {
	fmt.Println(a, b)
	degstir()
	fmt.Println(a, b)
}

/* aynı fonksıyonu global degısken kullanmadan tekrar yazalım .
package main
import"fmt"
func degıstır(x,y int)(int, int) {
return y,x
}
func main(){
// Yerel degıskenler fonksıyon ıcınde tanımlanır.
a:=1
b:=2
fmt.Println(a,b)
a,b= degistir(a,b)
fmt.Println(a,b)
}
*/
