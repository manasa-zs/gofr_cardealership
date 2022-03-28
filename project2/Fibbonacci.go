package project2

import "fmt"

func Fibbonacci() {
	a := 0
	b := 1
	for i := 1; i <= 10; i++ {
		c := a + b
		a = b
		b = c
		fmt.Println(c)
	}
}
