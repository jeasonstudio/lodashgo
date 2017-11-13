package main

import (
	"errors"
	"lodashgo/Array"
	"fmt"
)

func main() {
	fmt.Println("Hi, guys!")
	x := []int {1,2,3,4}
	a, _ := Array.IntReduce(x, reduceInt)
	y := []string {"a", "b", "c", "d"}
	b, _ := Array.StringReduce(y, reduceString)
	fmt.Println(a)
	fmt.Println(b)
	return
}

func reduceInt(accumulator interface{}, currentValue, currentIndex int, array []int) interface{} {
	switch accumulator.(type) {
		case int:
			return accumulator.(int) + currentValue
		default:
			return errors.New("")
	}
}

func reduceString(accumulator interface{}, currentValue string, currentIndex int, array []string) interface{} {
	switch accumulator.(type) {
		case string:
			return accumulator.(string) + currentValue
		default:
			return errors.New("")
	}
}