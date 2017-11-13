package main

import (
	"lodashgo/String"
	"fmt"
)

func main() {
	fmt.Println("Hi, guys!")
	r := String.EndsWith("aaabc", "ca")
	fmt.Println(r)
	return
}
