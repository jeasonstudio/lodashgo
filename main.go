package main

import (
	"lodashgo/String"
	"fmt"
)

func main() {
	fmt.Println("Hi, guys!")
	a := String.PadStart("11223", 15, "abc")
	fmt.Println(a)
	return
}
