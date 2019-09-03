package main

/*bir fonksıyonda panic olusursa, fonksıyonun defer kısmı ısletılıp fonksıyon durur.bu fonksıyonu
cagıran kısım kaldıgı yerden calısmaya devam eder.*/
import "fmt"

func kare(x int) {
	panic("durma noktasi") // buraya kadar saglam mı?
	fmt.Println(x * x)
}
func main() {
	kare(5)
}
