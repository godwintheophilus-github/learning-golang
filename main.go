package main

import "fmt"

func main() {
	fmt.Println("Welcome to Golang Learning")

	//declaration of variables in golang

	var x = 10
	y := 10
	var a int = 10
	var b int64
	b = 32000000000000000
	var c float64 = 3.14
	//var d int16 = 10043124321432 //./main.go:16:6: constant 10043124321432 overflows int16
	//fmt.Println(d)

	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(y)
	fmt.Println(a)
	fmt.Println(x)
}
