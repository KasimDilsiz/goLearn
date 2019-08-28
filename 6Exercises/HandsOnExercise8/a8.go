package main

import "fmt"

func main() {
x :=4
	switch  {
	case x<3:
		fmt.Println("kıs mevsimi" )
	case x>3:
		fmt.Println("ilkbahar mevsimi" )
	default:
		fmt.Println("iki mevsim de degil")
	}

}
