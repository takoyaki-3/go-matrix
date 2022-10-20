package main

import (
	"fmt"
	mat "gomatrix"
)

func main(){
	fmt.Println("hello")

	a := mat.NewMatrix(10,5)
	a.ScalarPlus(1)


	a.Sigmoid()
	a.Print()

	b := mat.NewRandom05to05Matrix(5,10)
	b.Sigmoid()
	b.Print()

	c:=mat.NewDotMatrix(a,b)
	c.Print()
	c.Sigmoid()
	c.Print()
}