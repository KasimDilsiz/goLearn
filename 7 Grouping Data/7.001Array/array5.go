package main

import "fmt"

func main() {
	nokta := [3][2]int{{11, 21}, {22, 32}, {33, 43}}
	for satir := 0; satir < 3; satir++ {
		for sutun := 0; sutun < 2; sutun++ {

			fmt.Print(nokta[satir][sutun], " ")
		}
		fmt.Println()
	}

}
