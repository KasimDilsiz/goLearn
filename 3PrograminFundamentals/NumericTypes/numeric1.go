package main

import "fmt"

/*uint8 ----> (0-255) 8-bit
uint16 -----> (0-65535) 16-bit
uint32 ----> (0-242949672) 32-bit
uint64 ----> (0-18446744073709551615) 64-bit

int8 ----> (-128 to 127) 8-bit
int16----> (-32768 to 32767) 16-bit
int32 ----> (...) 32-bit
int64 ----> (...) 64-bit
*/
func main() {
	var x int8 = 127
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

/* fakat bu kodumuz çalışmaz
var x int8=128
fmt.Println(x)
fmt.Printf("%T\n",x)

*/
