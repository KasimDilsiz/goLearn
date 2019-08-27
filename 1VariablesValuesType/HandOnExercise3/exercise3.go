package main

import "fmt"

var x int = 42
var y string = "Kasim Dilsiz"
var z bool = true

func main() {

	/*	fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)

	*/
	//BİR DİZGİYE YAZDIRMAK VE DAHA SONRA BU DEGERI BIR DEGISKENE VERMEK ISTERSEM FMT PAKETINDEN FUNC SPRINTF KULLANABILIRIM
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
