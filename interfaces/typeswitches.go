package main

import "fmt"

func typeSwitches(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("type of interface value is int")
	case string:
		fmt.Println("type of interface value is string")
	case float64:
		fmt.Println("type of interface value is float")
	default:
		fmt.Println("type is unknown")
	}
}

func main() {
	typeSwitches(10)
	typeSwitches("hello")
	typeSwitches(7.34)
}
