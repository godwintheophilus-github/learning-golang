package main

import "fmt"

func main() {
	var i interface{} = 6

	s, ok := i.(string)
	fmt.Println(s)
	fmt.Println(ok)

	s1, ok1 := i.(int)
	fmt.Println(s1)
	fmt.Println(ok1)
}
