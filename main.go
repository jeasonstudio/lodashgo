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
	t1 := []uint64 {}
	t2 := []string {}
	test := []interface{} {t1, t2}
	for _, tests := range test {
		switch t := tests.(type) {
			default:
				fmt.Printf("%T\n", t)
		}
	}
	return
}
