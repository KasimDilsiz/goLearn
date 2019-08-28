package main

import "fmt"


// mantıksal(boolen) bool ıfadesıyle tanımlanır//
func main()  {
	var x,y,z,t bool
	x=true ; y=false ;z=false;t=true
	fmt.Println(x&&y)
	fmt.Println(y||t)
	fmt.Println(!z)
	fmt.Println(x==y)

}