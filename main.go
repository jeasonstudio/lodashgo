package main

import "fmt"

func main() {
	fmt.Println("Hi, guys!")
	var map1 map[string]interface{}
	map1["shit"] = func() {
		fmt.Println("shit")
	}
	fmt.Println(map1)
	switch map1["shit"].(type) {
	case func():
		map1["shit"].(func())()
	}
	return
}
