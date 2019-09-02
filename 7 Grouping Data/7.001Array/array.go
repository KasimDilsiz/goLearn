package main

import "fmt"

/*Diziler,aynı türden veriden belirli sayida bulunduran degisken türüdür.
Dizi elemanlarina dogrudan indis numaralariyla erisilebilir.Dizi icerisinde farkli türden veri tutulamaz.
*/
func main() {
	var il [82] string
	il[1] = "adana"
	il[13] = "Bitlis"
	il[34] = "İstanbul"
	fmt.Println(il)
	fmt.Println(il[13])

}
