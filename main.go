package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Welcome to Golang Learning")
	var x string = "Hello"
	fmt.Println(x)

	//to get the length of a string
	fmt.Println(len(x)) //this will not give you the actual length of the string
	fmt.Println(utf8.RuneCountInString(x))
	fmt.Println(utf8.RuneCount([]byte(x)))
	fmt.Println([]byte(x))
}
