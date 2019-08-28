package main

import "fmt"

func main() {
	kd := 1998
	for {
		if kd > 2019 {
			break
		}
		fmt.Println(kd)
		kd++
	}
}
