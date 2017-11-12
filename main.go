package main

import (
	"lodashgo/Array"
	"fmt"
)

func main()  {
	fmt.Println("Hi, guys!")
	a := [] int {2}
	b := [] string {"s", "aa"}
	re, _ := Array.Concat(a, b, 3, "sss", true)
	// ree, _ := Array.Chunk(b, 2)
	fmt.Println(re)
	return
}
