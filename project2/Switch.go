package project2

import "fmt"

func Switch(s string) (string, int) {
	switch s {
	case "January", "December":
		fmt.Println("hi")
	case "Feb", "March":
		fmt.Println("hello")
	default:
		fmt.Println("no")
	}
	return "a", 1
}
