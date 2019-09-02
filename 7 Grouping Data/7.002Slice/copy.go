package main

/* copy fonksıyonunun avantajı; iki kesitin de eleman sayılarını kontrol etmesi ve en az olan eleman sayısı kadar elemanı
kopyalamasıdır.dolayısıyla kesıtlerin boyutlarından kaynaklanabılecek problemler engellenmıs olur .
*/
import (
	"fmt"
)

func main() {
	kaynak := []int{1, 2, 9, 9}
	hedef := make([]int, 2)
	fmt.Println(kaynak, hedef)
	sayi := copy(hedef, kaynak)
	fmt.Println("kopyalanan eleman sayısı:", sayi)
	fmt.Println(kaynak, hedef)
}
